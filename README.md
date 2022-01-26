# action-cppcheck

run cppcheck and post review comment to pull request

## Inputs

```yaml
inputs:
  github_token:
    description: "action github token"
    required: true
  repository:
    description: "owner and repository name"
    required: true
  pull_request_id:
    description: "pull request id"
    required: true
  allow_approve:
    description: "allow submit approve review"
    required: true
    default: true
```

## Allow approval

See: https://github.blog/changelog/2022-01-14-github-actions-prevent-github-actions-from-approving-pull-requests/
