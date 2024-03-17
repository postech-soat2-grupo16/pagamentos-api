provider "aws" {
  region = var.aws_region
}

#Configuração do Terraform State
terraform {
  backend "s3" {
    bucket = "terraform-state-soat"
    key    = "pagamentos-api/terraform.tfstate"
    region = "us-east-1"

    dynamodb_table = "terraform-state-soat-locking"
    encrypt        = true
  }
}

### Target Group + Listener

resource "aws_lb_target_group" "tg_pagamentos_api" {
  name        = "target-group-pagamentos-api"
  port        = 8000
  protocol    = "HTTP"
  target_type = "ip"
  vpc_id      = var.vpc_id

  health_check {
    enabled             = true
    interval            = 30
    matcher             = "200-299"
    path                = "/pagamentos/health"
    port                = "traffic-port"
    protocol            = "HTTP"
    timeout             = 5
    healthy_threshold   = 5
    unhealthy_threshold = 2
  }

  tags = {
    infra   = "target-group-pagamentos-api"
    service = "pagamentos"
  }
}

# Listener Rule that forwards the request to pagamentos-api TG
resource "aws_lb_listener_rule" "listener_pagamentos_api" {
  listener_arn = var.alb_fastfood_listener_arn
  priority     = 400

  condition {
    path_pattern {
      values = ["/pagamentos*"]
    }
  }

  action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.tg_pagamentos_api.arn
  }

  tags = {
    Name    = "alb-listener-pagamentos"
    infra   = "alb-listener-pagamentos"
    service = "pagamentos"
  }
}


### Task Config ###
resource "aws_ecs_task_definition" "task_definition_pagamentos_api" {
  family                   = "task-definition-pagamentos-api"
  requires_compatibilities = ["FARGATE"]
  network_mode             = "awsvpc"
  execution_role_arn       = var.execution_role_ecs
  task_role_arn            = var.execution_role_ecs

  cpu    = 512
  memory = 1024

  container_definitions = jsonencode([
    {
      name      = "container-pagamentos-api"
      image     = var.ecr_image
      cpu       = 512,
      memory    = 1024,
      essential = true,
      portMappings = [
        {
          containerPort = 8000
          hostPort      = 8000
          protocol      = "tcp"
          appProtocol   = "http"
        }
      ]

      environment = [
        { "name" : "DATABASE_URL", "value" : var.db_url },
        { "name" : "QUEUE_URL", "value" : var.sqs_url },
        { "name" : "IS_LOCAL", "value" : "false" },
        { "name": "NOTIFICATION_TOPIC", "value" : var.sns_arn}
      ]


      logConfiguration = {
        logDriver = "awslogs",
        options = {
          awslogs-create-group  = "true",
          awslogs-group         = "pagamentos-api-container",
          awslogs-region        = "us-east-1",
          awslogs-stream-prefix = "pagamentos"
        }
      },
    }
  ])

  runtime_platform {
    operating_system_family = "LINUX"
    cpu_architecture        = "X86_64"
  }

  tags = {
    infra    = "task-definition-pagamentos"
    services = "pagamentos"
  }
}

resource "aws_ecs_service" "ecs_service_pagamentos_api" {
  name                              = "ecs-service-pagamentos-api"
  cluster                           = var.ecs_cluster
  task_definition                   = aws_ecs_task_definition.task_definition_pagamentos_api.id
  launch_type                       = "FARGATE"
  platform_version                  = "1.4.0"
  desired_count                     = var.desired_tasks
  health_check_grace_period_seconds = 30

  network_configuration {
    subnets = [
      var.subnet_a,
      var.subnet_b
    ]
    security_groups  = [var.sg_cluster_ecs]
    assign_public_ip = false
  }

  load_balancer {
    target_group_arn = aws_lb_target_group.tg_pagamentos_api.arn
    container_name   = "container-pagamentos-api"
    container_port   = 8000
  }

  tags = {
    infra    = "ecs-service-pagamentos"
    services = "pagamentos"
  }
}