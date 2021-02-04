# Naga

![](https://www.ecured.cu/images/thumb/0/0f/Lady_vashj-wow.jpg/260px-Lady_vashj-wow.jpg)

Es un wrapper de viper que permite configurar una lista de variables de entrada, estableciendo la siguiente prioridad:

*  Flags
*  Variables de entorno
*  Archivos .yaml

Esta libreria recibe un **string** con el nombre del archivo yaml y una lista de estructuras de tipo ConfigEntry.

```
type ConfigEntry struct {
    VariableName string
    Description  string
    Shortcut     string
    DefaultValue interface{}
}
```

Esta estructura contiene los datos de las variables de configuración y se modifica cada uno de los DefaultValue según la configuración que se utilice. 

Una vez configurado, se retorna un map[string]interface{} con las llaves/valores configurados y un error. 

## Cómo usar

Primero descargamos la libreria y sus dependencias utilizando el comando:


```

go get "gitlab.falabella.com/fif/integracion/forthehorde/commons/naga"


```

Para utilizar naga se debe importar la librería:

```
import (
    "gitlab.falabella.com/fif/integracion/forthehorde/commons/naga"
)

```


Luego crear un VariableTypeResolver y un FlagConfigurator:

```
typeResolver := naga.NewVariableTypeResolver()
```

```
flagConfigurator := naga.NewFlagConfigurator(typeResolver)
```

Una vez creadas las estructuras, creamos un Configurator y lo ejecutamos.

```
configurator := naga.NewConfigurator(flagConfigurator, typeResolver)
values, err := configurator.Configure("filename", entries)
```
En donde *filename* es el nombre del archivo, en caso de no existir, ingresar un string vacío. Y *entries* es la lista de ConfigEntry, cómo por ejemplo:
```
entries := []naga.ConfigEntry{
		naga.ConfigEntry{
			VariableName: "example_one",
			Description:  "first entry",
			Shortcut:     "fe",
			DefaultValue: 0,
		},
		naga.ConfigEntry{
			VariableName: "example_two",
			Description:  "second entry",
			Shortcut:     "se",
			DefaultValue: "",
		},
		naga.ConfigEntry{
			VariableName: "example_three",
			Description:  "thrid entry",
			Shortcut:     "te",
			DefaultValue: false,
		},
	}
```




