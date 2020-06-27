const PROTO_PATH = __dirname + './../Server/proto/password.proto';
const grpc = require('grpc');
const protoLoader = require('@grpc/proto-loader');
// Suggested options for similarity to existing grpc.load behavior
const packageDefinition = protoLoader.loadSync(
    PROTO_PATH,
    {
        keepCase: true,
        longs: String,
        enums: String,
        defaults: true,
        oneofs: true
    });
const protoDescriptor = grpc.loadPackageDefinition(packageDefinition).proto;
// The protoDescriptor object has the full package hierarchy
const client = new protoDescriptor.PasswordGenerator('localhost:9090', grpc.credentials.createInsecure());

async function generate(length) {
    try {
        let call = await ((length) => {
            return new Promise((resolve, reject) => {
                client.generate(length, (error, response) => {
                    if (error) { reject(error); }
                    resolve(response);
                });
            });
        })({length});
        console.log(call);
    } catch (e) {
        console.log(`Ошибка: ${e.details}`);
    }
}

function main() {
    const len = process.argv[2];
    generate(len);
}

if (require.main === module) {
    main();
}

exports.generate = generate;
