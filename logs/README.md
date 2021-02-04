# logs

<Describir brevemente el proyecto>

## Diagrama de flujo

<Agregar aqui un diagrama de secuencia o de flujo describiendo la logica de la API>

## Como compilar

Para compilar simplemente se debe hacer `make build` lo que generara un binario en `build/bin/logs`

## Correr tests y coverage

Para correr los tests se debe hacer `make test` y para obtener el coverage `make coverage`

## Configuración

| Variables       |  Command && Shortcut   |                Descripción                |
| --------------- | :--------------------: | :---------------------------------------: |
| PORT            |      --port=, -p       | Puerto que escucha el servicio http, por defecto `8080`         |
| LOGGING_LEVEL   | --logging_level=,-l    | Nivel de logger, por defecto `info`                             |
| TRACING_ENABLED | --tracing_enabled=, -t | Indica si tracing se encuentra habilitado, por defecto `false`  |
| METRICS_ENABLED | --metrics_enabled=, -m | Indica si metrics se encuentra habilitado, por defecto `false`  |

## Ambientes

| Entorno |                 Url                  |
| ------- | :----------------------------------: |
| Dev     | http://integracion-k8s-dev.fif.tech/ |
| QA      | http://integracion-k8s-qa.fif.tech/  |
| PROD    | http://integracion-k8s-dcc.fif.tech/ |

## Anexos

<Agregar aqui referencias a documentacion externa>