# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: wikipage.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='wikipage.proto',
  package='protobuf',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=_b('\n\x0ewikipage.proto\x12\x08protobuf\"%\n\x04Wiki\x12\x1d\n\x05pages\x18\x01 \x03(\x0b\x32\x0e.protobuf.Page\"o\n\x04Page\x12\n\n\x02id\x18\x01 \x01(\x04\x12\r\n\x05title\x18\x02 \x01(\t\x12\r\n\x05links\x18\x03 \x03(\t\x12\x12\n\ncategories\x18\x04 \x03(\t\x12\x11\n\ttemplates\x18\x05 \x03(\t\x12\x16\n\x0enamedEntityIDs\x18\x06 \x03(\r\"4\n\x04\x44\x61ta\x12,\n\rnamedEntities\x18\x01 \x03(\x0b\x32\x15.protobuf.NamedEntity\"9\n\x0bNamedEntity\x12\n\n\x02id\x18\x01 \x01(\r\x12\x0e\n\x06jpName\x18\x02 \x01(\t\x12\x0e\n\x06\x65nName\x18\x03 \x01(\tb\x06proto3')
)




_WIKI = _descriptor.Descriptor(
  name='Wiki',
  full_name='protobuf.Wiki',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='pages', full_name='protobuf.Wiki.pages', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=28,
  serialized_end=65,
)


_PAGE = _descriptor.Descriptor(
  name='Page',
  full_name='protobuf.Page',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='protobuf.Page.id', index=0,
      number=1, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='title', full_name='protobuf.Page.title', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='links', full_name='protobuf.Page.links', index=2,
      number=3, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='categories', full_name='protobuf.Page.categories', index=3,
      number=4, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='templates', full_name='protobuf.Page.templates', index=4,
      number=5, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='namedEntityIDs', full_name='protobuf.Page.namedEntityIDs', index=5,
      number=6, type=13, cpp_type=3, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=67,
  serialized_end=178,
)


_DATA = _descriptor.Descriptor(
  name='Data',
  full_name='protobuf.Data',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='namedEntities', full_name='protobuf.Data.namedEntities', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=180,
  serialized_end=232,
)


_NAMEDENTITY = _descriptor.Descriptor(
  name='NamedEntity',
  full_name='protobuf.NamedEntity',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='protobuf.NamedEntity.id', index=0,
      number=1, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='jpName', full_name='protobuf.NamedEntity.jpName', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='enName', full_name='protobuf.NamedEntity.enName', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=234,
  serialized_end=291,
)

_WIKI.fields_by_name['pages'].message_type = _PAGE
_DATA.fields_by_name['namedEntities'].message_type = _NAMEDENTITY
DESCRIPTOR.message_types_by_name['Wiki'] = _WIKI
DESCRIPTOR.message_types_by_name['Page'] = _PAGE
DESCRIPTOR.message_types_by_name['Data'] = _DATA
DESCRIPTOR.message_types_by_name['NamedEntity'] = _NAMEDENTITY
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Wiki = _reflection.GeneratedProtocolMessageType('Wiki', (_message.Message,), dict(
  DESCRIPTOR = _WIKI,
  __module__ = 'wikipage_pb2'
  # @@protoc_insertion_point(class_scope:protobuf.Wiki)
  ))
_sym_db.RegisterMessage(Wiki)

Page = _reflection.GeneratedProtocolMessageType('Page', (_message.Message,), dict(
  DESCRIPTOR = _PAGE,
  __module__ = 'wikipage_pb2'
  # @@protoc_insertion_point(class_scope:protobuf.Page)
  ))
_sym_db.RegisterMessage(Page)

Data = _reflection.GeneratedProtocolMessageType('Data', (_message.Message,), dict(
  DESCRIPTOR = _DATA,
  __module__ = 'wikipage_pb2'
  # @@protoc_insertion_point(class_scope:protobuf.Data)
  ))
_sym_db.RegisterMessage(Data)

NamedEntity = _reflection.GeneratedProtocolMessageType('NamedEntity', (_message.Message,), dict(
  DESCRIPTOR = _NAMEDENTITY,
  __module__ = 'wikipage_pb2'
  # @@protoc_insertion_point(class_scope:protobuf.NamedEntity)
  ))
_sym_db.RegisterMessage(NamedEntity)


# @@protoc_insertion_point(module_scope)
