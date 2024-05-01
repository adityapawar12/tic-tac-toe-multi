from flask import Flask, jsonify, request
from sample.model.sample import Sample, SampleSchema

app = Flask(__name__)

samples = [
    Sample(1, 'test 1'),
    Sample(2, 'test 2'),
    Sample(3, 'test 3'),
    Sample(4, 'test 4'),
]

@app.route('/api/v1/sample')
def get_samples():
    schema = SampleSchema(many=True)
    return jsonify(schema.dump(samples))

@app.route('/api/v1/sample', methods=['POST'])
def add_sample():
    new_sample = SampleSchema().load(request.get_json())
    new_sample.id = len(samples) + 1
    samples.append(new_sample)
    return 'Created', 204

if __name__ == "__main__":
    app.run(port=3000)
