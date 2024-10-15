const net = require('net');
process.stdin.setEncoding('utf-8');

const client = new net.Socket();

client.connect(8080, '127.0.0.1', () => {
    console.log('Conectado al servidor TCP');
    // process.argv[2] Es el tercer argumento pasado al ejecutar el script
    // Por ejemplo, si ejecutamos "node index.js nombreUsuario", process.argv[2] sería "nombreUsuario"
    const nombreUsuario = process.argv[2] || 'UsuarioSinNombre';
    client.write("field//connectfield//" + nombreUsuario);
    process.stdout.write('Ingrese mensaje: ');
});

client.on('data', (data) => {
    console.log( `\n`+ data.toString());
    process.stdout.write('Ingrese mensaje: ');
});

client.on('close', () => {
    console.log('Conexión cerrada');
});

client.on('error', (err) => {
    console.error('Error: ', err.message);
});

process.stdin.on('data', (data) => {
    const mensaje = data.toString().trim(); 
    const msgToSend = `field//messagefield//${process.argv[2]}field//${process.argv[3]}field//`+mensaje;
    if (mensaje.toLowerCase() === 'exit') {
        client.end(); 
    } else {
        client.write(msgToSend); 
    }
});
