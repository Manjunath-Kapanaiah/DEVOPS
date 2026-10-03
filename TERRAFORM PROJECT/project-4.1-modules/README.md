# Project 4.1 — Terraform Modules, Outputs & State

Demonstrates reusable **modules**, **output variables**, and the **state file**.

## What it builds
Two reusable modules (`../modules/gcs-bucket`, `../modules/network`) are called from
`main.tf` to create a GCS bucket and a VPC + subnet.

## Feature status
| Feature | Status | Notes |
|---------|--------|-------|
| Reusable modules + outputs | **Implemented** | `gcs-bucket`, `network` |
| Remote GCS state backend | **Implemented** | `backend-gcs.tf` (bucket `likith-tfstate-capstone`, prefix `terraform/state`) |
| Uniform access + public-access prevention | **Implemented** | enforced in `modules/gcs-bucket` |
| Object lifecycle rule | Optional | enable via `lifecycle_age_days > 0` |
| CMEK encryption | Optional | enable via `kms_key_name` |

## Run it
```bash
gcloud auth application-default login          # one-time auth
cp terraform.tfvars.example terraform.tfvars   # then edit project_id

terraform init          # Task: initialise, download the google provider
terraform plan          # preview
terraform apply         # Task 5: create the infrastructure

terraform output        # Task 4: see the exposed output values
```

## Inspect the state (Task 6)
```bash
terraform state list                 # every resource Terraform tracks
terraform show                       # full state, human-readable
cat terraform.tfstate | less         # the raw JSON state file
```

## Modify & re-apply (Task 7)
Edit something (e.g. add a label in `main.tf`), then:
```bash
terraform plan     # Terraform shows the diff it detected
terraform apply    # it updates only what changed, and updates the state
```

## Deliverable screenshots
1. Project structure (`tree` or file explorer)
2. Module configuration (`main.tf` + `modules/`)
3. `terraform output` values
4. `terraform state list` / the state file
5. `terraform apply` after a config change (the diff)

## Elevate
- Remote GCS state is already active via `backend-gcs.tf`. Bootstrap the bucket once,
  then migrate local state:
  ```bash
  gcloud storage buckets create gs://likith-tfstate-capstone --location=US --uniform-bucket-level-access
  gcloud storage buckets update gs://likith-tfstate-capstone --versioning
  terraform init -migrate-state
  ```
- Add more modules / variables / outputs.

## Clean up
```bash
terraform destroy
```
