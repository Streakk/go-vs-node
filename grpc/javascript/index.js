import { loadPackageDefinition, Server, ServerCredentials } from '@grpc/grpc-js';
import { loadSync } from '@grpc/proto-loader';

// Load the .proto file
const packageDefinition = loadSync('../../proto/profile.proto', {
    keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true
});

const profileProto = loadPackageDefinition(packageDefinition).profile;

// Implement the GetUser RPC method
const getUser = (call, callback) => {
    callback(null, {
        userID: call.request.userID,
        username: 'JohnDoe',
        email: 'john.doe@example.com'
    });
};

// Create the gRPC server
const server = new Server();

// Bind the server to the implemented method
server.addService(profileProto.ProfileService.service, { GetUser: getUser });

// Start the server
server.bindAsync('127.0.0.1:50051', ServerCredentials.createInsecure(), () => {
    server.start();
    console.log('gRPC Node.js server running on http://127.0.0.1:50051');
});
