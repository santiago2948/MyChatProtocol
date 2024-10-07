const net = require('net');
process.stdin.setEncoding('utf-8');

// Crear un socket de cliente
const client = new net.Socket();

// Conectarse al servidor TCP en el puerto 8080
client.connect(8080, '127.0.0.1', () => {
    console.log('Conectado al servidor TCP');
    process.stdout.write('Ingrese mensaje: ');
});

// Escuchar mensajes del servidor
client.on('data', (data) => {
    console.log('\nMensaje del servidor: ' + data.toString());
    process.stdout.write('Ingrese mensaje: ');
});

// Manejar cierre de conexión
client.on('close', () => {
    console.log('Conexión cerrada');
});

// Manejar errores
client.on('error', (err) => {
    console.error('Error: ', err.message);
});

// Leer la entrada del usuario desde la consola
process.stdin.on('data', (data) => {
    const mensaje = data.toString().trim(); // Capturar entrada y eliminar espacios
    if (mensaje.toLowerCase() === 'exit') {
        client.end(); // Finaliza la conexión si el usuario escribe 'exit'
    } else {
        client.write(mensaje); // Enviar el mensaje al servidor
    }
});
