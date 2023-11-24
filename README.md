# action-workplace-notify

A GitHub Action to send a Workplace notification from your workflows.

## Usage

```yaml
- name: Notify
  uses: rotemtam/workplace-notify-action@v0
  with:
    access-token: ${{ secrets.WORKPLACE_ACCESS_TOKEN }}
    group-id: '1234567890'
    message: |
      # A markdown formatted message
      With [a link](https://rotemtam.com) 
```
 
## Access Token

1. Go to `https://<you org>.workplace.com/work/admin/apps/`
2. Click on `Create Custom Integration`.
3. Give it a name and description and click `Create`.
4. In the main screen click on `Create Access Token`.
5. Copy the token and save it as a secret in your repository.

## Permissions

The access token needs to have the following permissions:
* `manage_group_content`

### Legal

This project is licensed under the [Apache License 2.0](LICENSE).