const express = require('express');

const sampleRouter = express.Router();

let sampleData = [];

sampleRouter.get('/', (req, res) => {
  res.json(sampleData);
});

sampleRouter.get('/:id', (req, res) => {
  const id = req.params.id - 1;
  res.json(sampleData[id]);
});

sampleRouter.post('/', (req, res) => {
  const newData = req.body;
  console.log(newData);
  sampleData.push(newData);
  res.status(201).json(newData);
});

sampleRouter.put('/:id', (req, res) => {
  const id = req.params.id - 1;
  const updatedData = req.body;
  sampleData[id] = updatedData;
  res.json(updatedData);
});

sampleRouter.delete('/:id', (req, res) => {
  const id = req.params.id - 1;
  sampleData.splice(id, 1);
  res.sendStatus(204);
});

module.exports = sampleRouter;
