const express = require('express');
const bodyParser = require('body-parser');
const cors = require('cors');
const sampleRoutesV1 = require('./routes/sampleRoutesV1');

const app = express();
const PORT = 3000;

app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true })); // Parse application/x-www-form-urlencoded
app.use(cors());

app.use('/api/v1/sample', sampleRoutesV1);

app.listen(PORT, () => {
  console.log(`Server is running on http://localhost:${PORT}`);
});

