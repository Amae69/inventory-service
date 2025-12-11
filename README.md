# Inventory Service

A lightweight Golang API that exposes `/items` and `/status` endpoints.

## CI/CD Pipeline

This repository uses a comprehensive CI/CD pipeline defined in `.github/workflows/pipeline.yml`. It leverages reusable workflows to ensure consistency and best practices.

### Stages

1.  **Matrix Tests**:
    *   Uses `your-org/platform-ci/.github/workflows/go-matrix-test.yml`.
    *   Runs tests across multiple Go versions (1.22, 1.23, 1.24, 1.25) and Operating Systems (Ubuntu, macOS, Windows).
    *   Ensures cross-platform compatibility.

2.  **Docker Build & Push**:
    *   Uses `your-org/platform-ci/.github/workflows/build-and-push-image.yml`.
    *   Builds the Docker image using the `Dockerfile`.
    *   Tags the image as `ghcr.io/<your-username>/inventory-service:<sha>`.
    *   Pushes the image to GitHub Container Registry (GHCR).

3.  **Deploy to Kubernetes**:
    *   Uses `your-org/platform-ci/.github/workflows/deploy-k8s.yml`.
    *   Deploys the new image to the `inventory-staging` namespace in the Kubernetes cluster.
    *   Updates the `inventory-service` deployment.
    *   **Condition**: Runs only on the `main` branch after a successful Docker build.

### Secrets

The following secrets are required for the pipeline to function:

*   `REGISTRY_USERNAME`: Docker registry username (GitHub username).
*   `REGISTRY_PASSWORD`: Docker registry password (PAT with package read/write permissions).
*   `KUBECONFIG_STAGING`: Kubeconfig content for the staging cluster.

### Branching Strategy

*   **Feature Branches**: Create branches like `feature/description` for new work.
*   **Pull Requests**: Open PRs to `main` for review.
*   **Branch Protection**:
    *   Direct pushes to `main` are disabled.
    *   PR reviews are required.
    *   All CI checks must pass before merging.

## Application

### Endpoints

*   `GET /status`: Returns "OK".
*   `GET /items`: Returns a JSON list of items.

### Running Locally

```bash
go run main.go
```

### Running Tests

```bash
go test ./...
```
## End ...