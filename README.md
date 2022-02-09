# action-cppcheck

Check pull request with cppcheck and post result to review comments.

![Screenshot](screenshot.png)

## Inputs

```yaml
inputs:
  app_id:
    description: "github app id"
    required: false
  installation_id:
    description: "github app installation id"
    required: false
  private_key:
    description: "github app private key"
    required: false
  github_token:
    description: "action github token"
    required: false
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

## Example

### Run on local repo with github token

```yaml
name: cppcheck
on:
  pull_request:
    types: [opened, synchronize]
jobs:
  cppchceck:
    name: cppcheck
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - uses: myml/action-cppcheck@main
        with:
          github_token: ${{ secrets.GITHUB_TOKEN }}
          repository: ${{ github.repository }}
          pull_request_id: ${{ github.event.number }}
```

### Run on forked repo with github app

See [permissions-for-the-github_token](https://docs.github.com/en/actions/security-guides/automatic-token-authentication#permissions-for-the-github_token)

```yaml
name: cppcheck
on:
  pull_request:
    types: [opened, synchronize]
jobs:
  cppchceck:
    name: cppcheck
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - uses: myml/action-cppcheck@main
        with:
          app_id: xxx
          installation_id: xxx
          private_key: ${{ secrets.xxxx }}
          repository: ${{ github.repository }}
          pull_request_id: ${{ github.event.number }}
```

## Allow approval

See: https://github.blog/changelog/2022-01-14-github-actions-prevent-github-actions-from-approving-pull-requests/
