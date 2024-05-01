import datetime as dt

from marshmallow import Schema, fields

class Sample(object):
    def __init__(self, id, text):
        self.id = id 
        self.text = text

    def __repr__(self):
        return '<Sample(name={self.text!r})>'.format(self=self)


class SampleSchema(Schema):
    id = fields.Number()
    text = fields.Str()
