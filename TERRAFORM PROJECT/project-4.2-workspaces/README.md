# Project 4.2 — Multiple Environments with Terraform Workspaces

One configuration, three isolated environments (dev / staging / prod), each with
its **own state file**, using `terraform workspace`.

## Feature status
| Feature | Status | Notes |
|---------|--------|-------|
| Per-workspace env settings | **Implemented** | `local.env_settings` keyed by workspace |
| Workspace-name validation | **Implemented** | `null_resource.validate_workspace` fails fast on an unknown workspace |
| Remote GCS state backend | **Implemented** | `backend-gcs.tf`, prefix `terraform/workspaces` (each workspace gets its own sub-path) |

> With the remote backend active, run `terraform init -migrate-state` once to move any
> local `terraform.tfstate.d/*` state into GCS.

## Run it
```bash
cp terraform.tfvars.example terraform.tfvars   # edit project_id
terraform init

# Task 3 — create the workspaces
terraform workspace new dev
terraform workspace new staging
terraform workspace new prod
terraform workspace list          # deliverable: workspace list

# Task 4 — deploy DEV
terraform workspace select dev
terraform plan                    # deliverable: plan output
terraform apply

# Task 5 — deploy STAGING
terraform workspace select staging
terraform apply

# Task 6 — deploy PRODUCTION
terraform workspace select prod
terraform apply
```

## Verify environment isolation (Task 8)
```bash
terraform workspace select dev     && terraform output bucket_name
terraform workspace select prod    && terraform output bucket_name
# different bucket names => separate infrastructure + separate state
terraform state list               # each workspace tracks only its own resources
```
State files live under `terraform.tfstate.d/<workspace>/` — proof of isolation.

## Deliverable screenshots
1. `terraform workspace list`
2. `terraform workspace select ...` (switching)
3. `terraform plan` output
4. Buckets in the GCP console named `-dev-`, `-staging-`, `-prod-`

## Clean up (per workspace)
```bash
for ws in dev staging prod; do terraform workspace select $ws && terraform destroy; done
```
