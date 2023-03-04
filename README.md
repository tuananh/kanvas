# kanvas

`kanvas` is a high-level, go-to tool for deploying your container-based applications along with the whole infrastructure.

Tired of maintaining adhoc scripts or proprietary CI pipelines?

`kanvas`'s sole goal is to never require you to write an adhoc script to deploy your application. You write a kanvas configuration file once, and use it for various use-cases listed below...

## Use-cases

**Self-service isolated preview environments**: `kanvas` is ideal for providing self-service preview or deveploment environment. That is, every developer in your company can just run `kanvas apply` to get his own development/preview environment that contains all the VPC-related stuffs, ArgoCD, Kubernetes clusters, loadbalancers, and so on.

**Automate production deployments**: `kanvas` scales. Once you've satisfied with local tests, generate a proprietary CI config like one for GitHub Actions. It is able to run multiple Docker and Terraform jobs as concurrent as possible according to the DAG of the comonents leveraing the CI-specific concurrency primitive, like workflow jobs for GitHub Actions.

**Standardize how you deploy applications**: The examples contained in the `kanvas` project mainly targets EKS on AWS with ArgoCD. However, you can add any Terraform configs for ECS-based deployments or even GCP, Azure deployments.

## How it works

 `kanvas` enables you to declare plan-and-apply style deployment workflows that is powered by `Docker` and `Terraform`. You no longer need to write adhoc scripts, CircleCI(future) or GitHub Actions workflows to automate deployments.

## Getting started

`kanvas` is easy to get started. You don't need to reimplement your Docker and Terraform based deployment workflows using a completely different config management systems like Pulumi, AWS CDK, and so on.

Add `kanvas.yaml` to your project and write paths to directories and declare dependencies among directories, and you are good to go.

Under the hood, it runs on vanilla `docker` and `terraform`. Anything that can be deployed using the two should benefit from using `kanvas`.

`kanvas` is easy to test. Add `kanvas.yaml` to your config repository that contains `Dockerfile` and Terraform projects. Run `kanvas diff` to preview diffs and then run `kanvas apply` to deploy to the test environment.

## Configuration syntax

```
components:
  product1:
    components:
      appimage:
        dir: /containerimages/app
        docker:
          image: "davinci-std/example:myownprefix-"
      base:
        dir: /tf2
        needs:
        - appimage
        terraform:
          target: null_resource.eks_cluster
          vars:
          - name: containerimage_name
            valueFrom: appimage.id
      argocd:
        dir: /tf2
        needs:
        - base
        terraform:
          target: aws_alb.argocd_api
          vars:
          - name: cluster_endpoint
            valueFrom: base.cluster_endpoint
          - name: cluster_token
            valueFrom: base.cluster_token
      argocd_resources:
        # only path relative to where the command has run is supported
        # maybe the top-project-level "dir" might be supported in the future
        # which in combination with the relative path support for sub-project dir might be handy for DRY
        dir: /tf2
        needs:
        - argocd
        terraform:
          target: argocd_application.kanvas
```

The directory structure can be:

- product1/
  - base
  - argocd
  - argocd-resources

or

- product1/
  - 01-base
  - 02-argocd
  - 03-argocd-resources

In the second form, `depends_on` values are induced from the two digits before `-`, so that the projects are planned and applied in the ascending order of the denoted number.

The above configuration implies that:

- `kanvas plan` runs `terraform plan` and store the plan files up until the first unapplied terraform projects
- `kanvas apply` runs `docker build` and `terraform apply` for all the terraform projects planned beforehand

### Advanced: Environments

You can optionally add an `environments` field for defining two or more
environments. Each environment can have "defaults" and "after" fields.

**The "defaults" field**:

Each environment can have a "defaults" field for setting default values used for every component defined in the kanvas.yaml file. This helps making your multi-environment configuration DRY.

Why?

Let's say you have two "terraform" components in the config file and you want to use the "production" terraform workspace for production deployments while using the "preview" workspace for preview environments.

Without environments, you need to have two kanvas.yaml files and two pipelines/workflows/jobs//etc to run kanvas for respectively environments.

With environments, you can literally specify which terraform is used across all the terraform components per environment:

`kanvas.yaml`:

```
environments:
  production:
    defaults:
      terraform:
        workspace: production
  preview:
    defaults:
      terraform:
        workspace: preview

components:
  infra:
    dir: tf
    terraform:
      target: null_resource.infra
  k8s:
    needs:
    - infra
    dir: tf
    terraform:
      target: null_resource.k8s
```

Compare this is wrinting and maintaining two kanvas.yaml files like:

`production/kanvas.yaml`:

```
components:
  infra:
    dir: tf
    terraform:
      target: null_resource.infra
      workspace: production
  k8s:
    needs:
    - infra
    dir: tf
    terraform:
      target: null_resource.k8s
      workspace: production
```

`preview/kanvas.yaml`:

```
components:
  infra:
    dir: tf
    terraform:
      target: null_resource.infra
      workspace: preview
  k8s:
    needs:
    - infra
    dir: tf
    terraform:
      target: null_resource.k8s
      workspace: preview
```

Although the total LOC does not differ much, the version with environments have two benefits.

First, you no longer need to repeat "workspace" fields for all the components. This is especially nice when there are more components than environments, which might be common.

Second, you no longer need to duplicate whole components across environments. The more components you have, the nicer it becomes.

**The "after" field**:

Each environment can be configured to be deployed only after another envionment(s).

This is useful when you want some environment to be planned and applied only after other environments are successfully applied.

An example use-case for this would be to ensure changes working on the preview environment before you apply the changes to the production environment.

To make an environment (say "production") depend on another (say "preview"), add an `after` field under the environment and specify dependent environment names.

```
environments:
  production:
    after:
    - preview
  preview: {}
```

**Example**:

A more complete example of the whole configuration which involves `environments`, `defaults`, and `after` are shown below for your reference.

```
environments:
  production:
    after:
    - preview
    defaults:
      terraform:
        # This sets terraform.workspace for all the terraform components,
        # "infra" and "k8s" in this example.
        workspace: prod
  preview:
    defaults:
      terraform:
        # This sets terraform.workspace for all the terraform components,
        # "infra" and "k8s" in this example.
        workspace: prev

components:
  image:
    dir: dockerfile
    docker:
      image: "davinci-std/example:myownprefix-"
    # We opt to use te specific image repository only for the preview environments:
      preview:
        docker:
          image: "myrepo/example:myownprefix-"
  infra:
    dir: tf
    terraform:
      target: null_resource.infra
  k8s:
    needs:
    - infra
    dir: tf
    terraform:
      target: null_resource.k8s
  app:
    needs:
    - k8s
    - image
```

## Roadmap

Here's the preliminary list of roadmap items to be implemented:

- [ ] Ability to specify the Terraform project templates for reuse
- [ ] Ability to export the workflow to CodeBuild (Multiple kanvas jobs are mapped to a single build)
- [ ] Ability to export the workflow to GitHub Actions (Each kanvas job is mapped to one Actions job)
- [ ] An example project that covers kompose, EKS, and ArgoCD.
  - We'll be using [vals](https://github.com/helmfile/vals) for embedding secret references in `docker-compose.yml`, [kompose](https://kompose.io/) for converting `docker-compose.yml` to Kubernetes manifests, apply or git-push the manifests somehow(perhaps we'll be creating a dedicated tool for that), and [terraform-argocd-provider](https://github.com/oboukili/terraform-provider-argocd) for managing ArgoCD projects, applications, cluster secrets, and so on.

## Related projects

`kanvas` is supposed to be a spritual successor to [Atlantis](https://github.com/runatlantis/atlantis). Like Atlantis, we needed something that is convenient to power mainly-Terraform-based workflows. We also needed out-of-box support for more tools, wanted it be easier to operate and easier to test locally and easier to integrate into various environments.

## References

We've carefully surveyed the current state of Terraform and its wrappers in the opensource community and found something like `kanvas` is missing and here we are.

- https://github.com/hashicorp/terraform/issues/25244
- https://github.com/hashicorp/terraform/issues/30937
- https://github.com/hashicorp/terraform/issues/4149
- https://discuss.hashicorp.com/t/multiple-plan-apply-stages/8320
- https://developer.hashicorp.com/terraform/language/state/remote-state-data
- https://registry.terraform.io/providers/hashicorp/tfe/latest/docs/data-sources/outputs
- https://terragrunt.gruntwork.io/docs/features/work-with-multiple-aws-accounts/ terragrunt is great but scales vertically
- https://github.com/transcend-io/terragrunt-atlantis-config generates atlantis config for terragrunt. We generate GitHub Actions workflow config for our own terraform+alpha wrapper
- https://terragrunt.gruntwork.io/docs/features/execute-terraform-commands-on-multiple-modules-at-once/#dependencies-between-modules terragrunt manages dependencies amongn terraform projects. We manage dependencies across any supported jobs (currently docker and terraform)

