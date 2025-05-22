import * as grpc from '@grpc/grpc-js';
import * as protoLoader from '@grpc/proto-loader';

const packageDef = protoLoader.loadSync('../../pkg/services/factors/factors.proto');
const proto = grpc.loadPackageDefinition(packageDef) as any;

enum OperationType {
    SUM = 'SUM',
    SUBTRACTION = 'SUBTRACTION',
    MULTIPLICATION = 'MULTIPLICATION',
    DIVISION = 'DIVISION'
}

enum SourceType {
    BONUS = 'BONUS',
    DISCOUNT = 'DISCOUNT',
}

const server = new grpc.Server();

server.addService(proto.factors.FactorsService.service, {
    GetFactor: (call: any, callback: any) => {
        const source = call.request.source;

        let factor = 0;
        switch (source) {
            case SourceType.BONUS:
                factor = 2;
                break;
            case SourceType.DISCOUNT:
                factor = 1;
                break;
        }
        callback(null, { value: factor.toFixed(2) });
    }
});

server.bindAsync('0.0.0.0:50052', grpc.ServerCredentials.createInsecure(), () => {
    console.log('FactorsService gRPC server running on :50052');
});