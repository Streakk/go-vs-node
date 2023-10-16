import { createServer } from 'http';
const PORT = 8080;

const server = createServer((req, res) => {
    if (req.method === 'POST' && req.url === '/compute') {
        let body = '';

        req.on('data', (chunk) => {
            body += chunk;
        });

        req.on('end', () => {
            try {
                const values = JSON.parse(body).values;

                if (!Array.isArray(values)) {
                    res.writeHead(400, { 'Content-Type': 'application/json' });
                    res.end(JSON.stringify({ error: 'Bad Request' }));
                    return;
                }

                const sum = values.reduce((acc, val) => acc + val, 0);
                const product = values.reduce((acc, val) => acc * val, 1);
                const average = sum / values.length;

                res.writeHead(200, { 'Content-Type': 'application/json' });
                res.end(JSON.stringify({
                    sum: sum,
                    average: average,
                    product: product
                }));
            } catch (err) {
                res.writeHead(400, { 'Content-Type': 'application/json' });
                res.end(JSON.stringify({ error: 'Bad Request' }));
            }
        });
    } else {
        res.writeHead(404, { 'Content-Type': 'application/json' });
        res.end(JSON.stringify({ error: 'Not Found' }));
    }
});

server.listen(PORT, () => {
    console.log(`Server is running on port ${PORT}`);
});
