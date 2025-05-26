# Argo Rollouts Automation Example

This repository demonstrates how to set up:

- A Go app with `/good` and `/bad` endpoints
- GitHub Actions to build + push Docker images on every `main` push
- ArgoCD + Argo Image Updater for GitOps deployment
- Argo Rollouts for progressive delivery
- Prometheus + Grafana to monitor and trigger automated rollback

## Setup Instructions

1. **Build the app**

```bash
docker build -t your-dockerhub/example-app:latest .
docker push your-dockerhub/example-app:latest
```

2. **Install ArgoCD and Argo Rollouts**

```bash
kubectl create namespace argocd
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
kubectl apply -f https://github.com/argoproj/argo-rollouts/releases/latest/download/install.yaml
```

3. **Set up ArgoCD Application**

Apply the manifests under `/manifests` to ArgoCD.

4. **Test the Deployment**

```bash
kubectl port-forward svc/example-app 8080:8080
curl http://localhost:8080/good
```

5. **Add the `/bad` endpoint and push changes**

Commit + push → GitHub Action triggers → Argo Image Updater updates deployment → Argo Rollouts deploys.

6. **Monitor with Prometheus + Grafana**

Set up Prometheus and Grafana using the provided configs under `/prometheus` and `/grafana`.

## Development

### Prerequisites

- Go 1.21 or later
- Docker
- kubectl
- [Argo Rollouts CLI](https://argoproj.github.io/argo-rollouts/installation/#kubectl-plugin)

### Local Development

1. **Clone and Install Dependencies**

```bash
git clone https://github.com/yourusername/argo-rollouts-automation-example.git
cd argo-rollouts-automation-example
cd app
go mod download
```

2. **Run the Application Locally**

```bash
go run main.go
```

The application will start on `http://localhost:8080` with the following endpoints:
- `GET /good` - Returns a 200 OK response
- `GET /bad` - Returns a 500 Internal Server Error

3. **Build and Test**

```bash
# Run tests
go test ./...

# Build the application
go build -o example-app

# Build Docker image locally
docker build -t example-app:dev .
```

4. **Development Workflow**

- Make changes to the application code
- Run tests locally
- Build and test the Docker image
- Push changes to a feature branch
- Create a pull request to `main`
- Once merged, the GitHub Action will build and push the new image
- Argo Image Updater will detect the new image and trigger a rollout

### Project Structure

```
.
├── app/
    └── main.go         # Application entry point
├── Dockerfile          # Container definition
├── manifests/          # Kubernetes manifests
│   ├── deployment.yaml
│   └── rollout.yaml
├── prometheus/         # Prometheus configuration
├── grafana/            # Grafana dashboards
└── .github/            # GitHub Actions workflows
```

## Example Commands to Test Phases

- Test success:
```bash
curl http://example-app/good
```

- Test failure:
```bash
curl http://example-app/bad
```

- Check rollout status:
```bash
kubectl argo rollouts get rollout example-app --watch
```

- Roll back manually if needed:
```bash
kubectl argo rollouts undo example-app
```

Happy deploying!

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
