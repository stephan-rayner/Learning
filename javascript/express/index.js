const express = require('express');
const app = express();
const bodyParser = require('body-parser');
app.use(bodyParser.json());

app.post('/event', (request, response) => {
  console.log('body', request.body.hello)
});

app.listen(3000, () => {
  console.log('Example app listening on port 3000!');
});
