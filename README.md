# Literature List Search

## Getting started
This is an example of how you may give instructions on setting up your project locally. To get a local copy up and running follow these example steps.

### Clone the repository
```shell
git clone https://github.com/zenorachi/literature-list
```
### [Optional] Setup configs
(_./configs/main.yml_)
```yaml
http:
  host: 0.0.0.0
  port: 8080
  readTimeout: 10s
  writeTimeout: 10s
  shutdownTimeout: 5s

cyberleninkaClient:
  baseUrl: "https://cyberleninka.ru/"
  timeout: 10s
```
> **Note:** everything is already configured by default
### Compile and run the project
```shell
make
```
### Example of request and response
```shell
curl --location --request GET 'http://127.0.0.1:8080/api/v1/literature/search' \
--header 'Content-Type: application/json' \
--data '{
  "literature_list": [
    "test"
  ]
}'
```
```json
{
    "literature_list": [
        {
            "title": "test",
            "is_contained": true
        }
    ]
}
```
### Stop the running container
```shell
make stop
```