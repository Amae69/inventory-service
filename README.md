# Inventory Service

A lightweight Golang API that exposes `/items` and `/status` endpoints.

## CI/CD Pipeline

This project uses GitHub Actions for Continuous Integration and Continuous Deployment. The pipeline is defined in `.github/workflows/pipeline.yml`.

### Workflow Structure

The pipeline consists of the following jobs:
1.  **setup**: Prepares the build environment and defines the image name.
2.  **tests**: Runs the Go test suite using a matrix strategy.
3.  **docker**: Builds and pushes the Docker image to GitHub Container Registry (GHCR).
4.  **deploy-staging**: Deploys the application to a Kubernetes cluster (staging environment).

### Reusable Workflows

We utilize reusable workflows from the `platform-ci` repository to maintain consistency and reduce duplication:
-   `go-matrix-test.yml`: Standardized Go testing workflow.
-   `build-and-push-image.yml`: Standardized Docker build and push workflow.
-   `deploy-k8s.yml`: Standardized Kubernetes deployment workflow.

### Matrix Tests

The `tests` job uses a matrix strategy to ensure compatibility across multiple environments:
-   **Go Versions**: 1.22, 1.23, 1.24, 1.25
-   **Operating Systems**: Ubuntu, macOS, Windows

This ensures that the application is robust and works correctly on different platforms and Go versions.

### Secrets Management

The pipeline relies on the following GitHub Secrets:
-   `REGISTRY_USERNAME` / `REGISTRY_PASSWORD`: Credentials for pushing images to GHCR.
-   `KUBECONFIG_STAGING`: Kubeconfig file content for accessing the staging cluster.

### Branching Strategy

-   **main**: The primary branch. Pushes to `main` trigger the full pipeline, including deployment to staging.
-   **feature/\*\***: Feature branches. Pushes trigger tests and build steps but skip deployment.
-   **Pull Requests**: PRs targeting `main` trigger the CI pipeline (tests) to ensure code quality before merging.

### Deployment Strategy

The deployment process involves:
1.  **Containerization**: The application is packaged as a Docker image.
2.  **Registry**: The image is pushed to GitHub Container Registry.
3.  **Kubernetes**: The application is deployed to a Kubernetes namespace (`inventory-staging`) using a standard Deployment resource.
    -   It creates/updates the necessary secrets for image pulling.
    -   It performs a rollout and waits for the deployment to be ready.

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
## End ..

