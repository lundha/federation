import { ApolloServer } from '@apollo/server';
import { startStandaloneServer } from '@apollo/server/standalone';
import { ApolloGateway, IntrospectAndCompose, RemoteGraphQLDataSource } from '@apollo/gateway';
import { serializeQueryPlan } from '@apollo/query-planner';

class DebugDataSource extends RemoteGraphQLDataSource {
    willSendRequest({
        request
    }) {
        if (request.operationName) {
            console.log(`Operation name: ${request.operationName}`);
        }
        console.log(`Query body: ${request.query}`);
    }
}

const gateway = new ApolloGateway({
    experimental_didResolveQueryPlan: opts => {
        if (opts.requestContext.operationName != "IntrospectionQuery") {
            // clear console
            console.clear();
            console.log(serializeQueryPlan(opts.queryPlan));
        }
    },
    supergraphSdl: new IntrospectAndCompose({
        subgraphs: [
            { name: 'contracts', url: 'http://localhost:4001/query' },
            { name: 'users', url: 'http://localhost:4002/query' },
            { name: 'suppliers', url: 'http://localhost:4003/query' }
        ],
    }),
    buildService({ url }) {
        return new DebugDataSource({ url });
    }
});

const server = new ApolloServer({
    gateway,
    subscriptions: false,
});

const requestTimePlugin = {
    requestDidStart() {
        const start = Date.now();
        let op = ""
        return {
            didResolveOperation({ operationName }) {
                if (operationName) {
                    op = operationName;
                }
            },
            willSendResponse() {
                if (op === "IntrospectionQuery") {
                    return;
                }
                const stop = Date.now();
                const elapsed = stop - start;
                console.log(`Operation ${op} took: ${elapsed}ms`);
            },
        };
    },
};

server.addPlugin(requestTimePlugin);

// Note the top-level `await`!
const { url } = await startStandaloneServer(server);
console.log(`🚀  Server ready at ${url}`);
