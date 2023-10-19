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

                let sum = 0;
                let product = 1;
                for (let i = 0; i < values.length; i++) {
                    sum += values[i];
                    product *= values[i];
                }                
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
