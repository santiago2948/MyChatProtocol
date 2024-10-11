# Proyecto de Chat Cliente-Servidor en Go y Node.js


## Content

- [Introducción](#introducción)
- [Objetivos](#objetivos)
- [Arquitectura](#arquitectura)
- [Requisitos](#requisitos)
- [Instalación](#instalación)
  - [Servidor](#servidor)
  - [Cliente](#cliente)
- [Desarrollo](#desarrollo)
  - [Cliente (`index.js`)](#cliente-indexjs)
  - [Servidor (`main.go`)](#servidor-maingo)
- [Comunicacion](#comunicacion)
- [Aspectos Logrados](#aspectos-logrados)
- [Aspectos No Logrados](#aspectos-no-logrados)
- [Conclusiones](#conclusiones)
- [Referencias](#referencias)

## Introducción

Este proyecto es un sistema de chat cliente-servidor que permite la comunicación en tiempo real entre varios usuarios a través de una arquitectura cliente-servidor utilizando sockets TCP. El cliente está desarrollado en **Node.js** y el servidor en **Go**, lo que permite demostrar la interoperabilidad de sistemas en red usando diferentes tecnologías.

## Objetivos

- Desarrollar una aplicación de chat cliente-servidor.
- Implementar un protocolo de capa de aplicación para el intercambio de mensajes.
- Soportar múltiples conexiones concurrentes.
- Utilizar sockets TCP para la comunicación entre el cliente y el servidor.

## Arquitectura

El proyecto sigue una arquitectura cliente-servidor donde:

- El **cliente** es responsable de conectarse al servidor, enviar mensajes y recibir mensajes de otros usuarios.
- El **servidor** gestiona múltiples conexiones de clientes, retransmite los mensajes a los destinatarios correctos y maneja la concurrencia entre los diferentes usuarios conectados.

## Requisitos

- **Cliente**: Node.js (versión >= 14.0.0)
- **Servidor**: Go (versión >= 1.23.1)
- Conexión a la red para comunicación entre el cliente y el servidor.

## Instalación

### Servidor

1. Clona el repositorio.
2. Navega al directorio del servidor: `cd server/`
3. Ejecuta `go mod tidy` para instalar las dependencias.
4. Ejecuta el servidor con el siguiente comando:
   ```bash
   go run main.go

### Cliente 

1. Clona el reprositorio.
2. Navega al directorio del cliente: `cd client/`.
3. Ejecuta `npm intall` para instalar las dependencias.
4. Ejecuta el cliente con el siguiente comando:
   ```bash
   node index.js [nombreUsuario]

## Desarrollo

### Cliente(`index.js`)

- El cliente se conecta al servidor TCP en el puerto `8080` y envía mensajes a otros usuarios.
- Usa la biblioteca `net` de Node.js para gestionar las conexiones TCP.
- Los mensajes siguen el formato:
  `field//messagefield//nombreUsuariofield//nombreDestinatariofield//mensaje`.
  #### Principales Funcionalidades:
  1. Conexión al servidor TCP en `127.0.0.1` y puerto `8080`.
  2. El nombre de usuario se toma como argumento al ejecutar el script (`node index.js nombreUsuario`).
  3. Se recibe y muestra en consola los mensajes enviados desde el servidor.
  4. El cliente envía mensajes con la estructura
    `field//messagefield//nombreUsuariofield//nombreDestinatariofield//mensaje`.

### Servidor(`main.go`)

- El servidor escucha conexiones TCP en el puerto `8080` y gestiona los mensajes entre los clientes.
- Utiliza un mapa para identificar a los clientes conectados a través de sus nombres de usuario.
- Los mensajes se procesan usando la estructura `Message` y se envían al destinatario adecuado.
  #### Principales Funcionalidades:
  1. Gestión de múltiples clientes mediante el uso de `goroutines`.
  2. El servidor almacena los clientes conectados en un mapa `clients`, asociando su `Nickname` con su conexión.
  3. Los mensajes entre clientes se transmiten utilizando la estructura `Message` definida en `message.go`.
  4. Cuando un cliente se desconecta, el servidor elimina su entrada del mapa y muestra un mensaje de desconexión.

## Comunicacion

Los mensajes entre el cliente y el servidor se estructuran de la siguiente manera:
- Conexión: field//connectfield//nombreUsuario
- Mensajes:
  `field//messagefield//nombreUsuariofield//nombreDestinatariofield//contenidoMensaje`

## Aspectos Logrados 

- Conexión exitosa entre cliente y servidor.
- Envío y recepción de mensajes entre varios clientes.
- Manejo de múltiples conexiones concurrentes.
- Mecanismo básico de desconexión de clientes.

## Aspectos No Logrados 

- No se implementó una interfaz gráfica para el cliente.
- No se manejan mensajes persistentes ni se guarda el historial de chat.

## Cloclusiones 

Este proyecto muestra una implementación básica de un sistema de chat cliente-servidor utilizando Go y Node.js, lo que demuestra el uso de sockets para comunicación en red. Se logró la interoperabilidad entre los sistemas cliente y servidor, y la gestión de múltiples conexiones simultáneas.

## Referencias

- [TCP Server-Client Implementation in Go](https://www.geeksforgeeks.org/tcp-server-client-implementation-in-c/) [1]
- [Beej's Guide to Network Programming](https://beej.us/guide/bgnet/) [1]
