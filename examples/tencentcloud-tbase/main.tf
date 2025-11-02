# TDSQL PostgreSQL (Tbase) Examples

# ========== Data Sources ==========

# Query Tbase instances
data "cloud_tbase_instances" "instances" {
  instance_id = "tbase-xxxxx"
}

# Query Tbase PG instances
data "cloud_tbase_pg_instances" "pg_instances" {
  instance_id = "tbase-pg-xxxxx"
}

# ========== Resources ==========

# Tbase Instance
resource "cloud_tbase_instance" "instance" {
  instance_name = "example-tbase"
  availability_zone = "ap-guangzhou-3"
  vpc_id        = "vpc-xxxxx"
  subnet_id     = "subnet-xxxxx"
  db_version    = "10.0"
  cpu           = 4
  memory        = 8
  storage       = 100
  
  tags = {
    env = "test"
  }
}

# Tbase PG Instance
resource "cloud_tbase_pg_instance" "pg_instance" {
  instance_name     = "example-tbase-pg"
  availability_zone = "ap-guangzhou-3"
  vpc_id            = "vpc-xxxxx"
  subnet_id         = "subnet-xxxxx"
  db_version        = "10.0"
  cpu               = 4
  memory            = 8
  storage           = 100
  
  tags = {
    env = "test"
  }
}

# Tbase PG Instance VIP
resource "cloud_tbase_pg_instance_vip" "vip" {
  instance_id = cloud_tbase_pg_instance.pg_instance.id
  vip         = "10.0.1.100"
  subnet_id   = "subnet-xxxxx"
}
