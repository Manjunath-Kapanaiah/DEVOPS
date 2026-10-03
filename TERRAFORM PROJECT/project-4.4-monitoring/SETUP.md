# Monitoring Stack — Setup on a Fresh Linux Machine

Runs Prometheus + Grafana + node-exporter + kube-state-metrics on a local
`kind` cluster via the `kube-prometheus-stack` Helm chart. No cloud cost.

## Prerequisites

- Linux (Ubuntu/Debian or WSL2)
- ~4 GB free RAM, internet access

---

## Step 0 — Install tooling (one-time per machine)

```bash
# Docker
curl -fsSL https://get.docker.com | sudo sh
sudo usermod -aG docker $USER
newgrp docker                     # or log out/in

# kubectl
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl && rm kubectl

# kind
curl -Lo ./kind https://kind.sigs.k8s.io/dl/v0.27.0/kind-linux-amd64
sudo install -o root -g root -m 0755 kind /usr/local/bin/kind && rm kind

# helm
curl -fsSL https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 | bash
```

Verify:

```bash
docker --version && kind --version && kubectl version --client && helm version
docker ps          # Docker daemon must be running (WSL2: sudo service docker start)
```

---

## Step 1 — Get the project

```bash
git clone <your-repo-url>
cd project-4.4-monitoring          # must contain kind-config.yaml and values.yaml
```

## Step 2 — Create the cluster

```bash
kind create cluster --name monitoring --config kind-config.yaml
kubectl config get-contexts        # * should be on kind-monitoring
kubectl get nodes                  # both nodes Ready (wait ~30s)
```

## Step 3 — Install the monitoring stack

```bash
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
kubectl create namespace monitoring
helm install kps prometheus-community/kube-prometheus-stack -n monitoring -f values.yaml
kubectl get pods -n monitoring -w  # Ctrl+C once all Running/Ready
```

📸 **Screenshot 1** — `kubectl get pods -n monitoring` (all Running).

## Step 4 — Open Grafana

```bash
kubectl port-forward -n monitoring svc/kps-grafana 3000:80
```

Browse http://localhost:3000 → login **admin / admin123**.

📸 **Screenshot 2** — Grafana login page + grafana pod Running.

## Step 5 — Verify Prometheus data source

Grafana → **Connections → Data sources → Prometheus** (auto-configured).

📸 **Screenshot 3** — the Prometheus data source.

## Step 6 — Import dashboard

Grafana → **Dashboards → New → Import → ID `1860` → Load → select Prometheus → Import**.

📸 **Screenshot 4** — dashboard 1860 with CPU/memory/pods.

---

## Teardown

```bash
kind delete cluster --name monitoring
```

---

## Troubleshooting: `kubectl get nodes` fails

This only happens when the kind containers aren't running (reboot / Docker restart).

```bash
docker ps                                                   # are the containers Up?
docker start monitoring-control-plane monitoring-worker     # if stopped, start them
kubectl config use-context kind-monitoring                  # if context is wrong
kind export kubeconfig --name monitoring                    # rebuild kubeconfig entry
```
