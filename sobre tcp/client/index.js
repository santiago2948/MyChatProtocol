const net = require('net');
process.stdin.setEncoding('utf-8');

const client = new net.Socket();

client.connect(8080, '127.0.0.1', () => {
    console.log('Conectado al servidor TCP'); 
    // process.argv[2] is the third argument passed when running the script
    // For example, if we run "node index.js username", process.argv[2] would be "username"
    const username = process.argv[2] || 'UsuarioSinNombre';
    client.write("field//connectfield//" + username);
    process.stdout.write('Ingrese mensaje: ');
});

client.on('data', (data) => {
    console.log('\nMensaje del servidor: ' + data.toString());
    process.stdout.write('Ingrese mensaje: ');
});

client.on('close', () => {
    console.log('Conexión cerrada');
});

client.on('error', (err) => {
    console.error('Error: ', err.message);
});

process.stdin.on('data', (data) => {
    const message = data.toString().trim(); 
    const msgToSend = `field//messagefield//${process.argv[2]}field//${process.argv[3]}field//` + message;
    if (message.toLowerCase() === 'exit') {
        client.end(); 
    } else {
        client.write(msgToSend); 
    }
});
