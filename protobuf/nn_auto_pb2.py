# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: nn_auto.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='nn_auto.proto',
  package='protobuf',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=_b('\n\rnn_auto.proto\x12\x08protobuf\"n\n\x04Wiki\x12(\n\x05pages\x18\x01 \x03(\x0b\x32\x19.protobuf.Wiki.PagesEntry\x1a<\n\nPagesEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\x1d\n\x05value\x18\x02 \x01(\x0b\x32\x0e.protobuf.Page:\x02\x38\x01\"\x82\x01\n\x04Page\x12\n\n\x02id\x18\x01 \x01(\x04\x12\r\n\x05title\x18\x02 \x01(\t\x12\x11\n\tnamespace\x18\x03 \x01(\r\x12\r\n\x05links\x18\x04 \x03(\t\x12\x12\n\ncategories\x18\x05 \x03(\t\x12\x11\n\ttemplates\x18\x06 \x03(\t\x12\x16\n\x0enamedEntityIDs\x18\x07 \x03(\r\"\x96\x02\n\x04\x44\x61ta\x12\x38\n\rnamedEntities\x18\x01 \x03(\x0b\x32!.protobuf.Data.NamedEntitiesEntry\x12\x34\n\x0btaggedPages\x18\x02 \x03(\x0b\x32\x1f.protobuf.Data.TaggedPagesEntry\x1aK\n\x12NamedEntitiesEntry\x12\x0b\n\x03key\x18\x01 \x01(\r\x12$\n\x05value\x18\x02 \x01(\x0b\x32\x15.protobuf.NamedEntity:\x02\x38\x01\x1aQ\n\x10TaggedPagesEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12,\n\x05value\x18\x02 \x01(\x0b\x32\x1d.protobuf.TaggedNamedEntities:\x02\x38\x01\"9\n\x0bNamedEntity\x12\n\n\x02id\x18\x01 \x01(\r\x12\x0e\n\x06jpName\x18\x02 \x01(\t\x12\x0e\n\x06\x65nName\x18\x03 \x01(\t\"$\n\x13TaggedNamedEntities\x12\r\n\x05\x61ttrs\x18\x01 \x03(\rb\x06proto3')
)




_WIKI_PAGESENTRY = _descriptor.Descriptor(
  name='PagesEntry',
  full_name='protobuf.Wiki.PagesEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='protobuf.Wiki.PagesEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='protobuf.Wiki.PagesEntry.value', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=_b('8\001'),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=77,
  serialized_end=137,
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
  nested_types=[_WIKI_PAGESENTRY, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=27,
  serialized_end=137,
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
      name='namespace', full_name='protobuf.Page.namespace', index=2,
      number=3, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='links', full_name='protobuf.Page.links', index=3,
      number=4, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='categories', full_name='protobuf.Page.categories', index=4,
      number=5, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='templates', full_name='protobuf.Page.templates', index=5,
      number=6, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='namedEntityIDs', full_name='protobuf.Page.namedEntityIDs', index=6,
      number=7, type=13, cpp_type=3, label=3,
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
  serialized_start=140,
  serialized_end=270,
)


_DATA_NAMEDENTITIESENTRY = _descriptor.Descriptor(
  name='NamedEntitiesEntry',
  full_name='protobuf.Data.NamedEntitiesEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='protobuf.Data.NamedEntitiesEntry.key', index=0,
      number=1, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='protobuf.Data.NamedEntitiesEntry.value', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=_b('8\001'),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=393,
  serialized_end=468,
)

_DATA_TAGGEDPAGESENTRY = _descriptor.Descriptor(
  name='TaggedPagesEntry',
  full_name='protobuf.Data.TaggedPagesEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='protobuf.Data.TaggedPagesEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='protobuf.Data.TaggedPagesEntry.value', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=_b('8\001'),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=470,
  serialized_end=551,
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
    _descriptor.FieldDescriptor(
      name='taggedPages', full_name='protobuf.Data.taggedPages', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_DATA_NAMEDENTITIESENTRY, _DATA_TAGGEDPAGESENTRY, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=273,
  serialized_end=551,
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
  serialized_start=553,
  serialized_end=610,
)


_TAGGEDNAMEDENTITIES = _descriptor.Descriptor(
  name='TaggedNamedEntities',
  full_name='protobuf.TaggedNamedEntities',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='attrs', full_name='protobuf.TaggedNamedEntities.attrs', index=0,
      number=1, type=13, cpp_type=3, label=3,
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
  serialized_start=612,
  serialized_end=648,
)

_WIKI_PAGESENTRY.fields_by_name['value'].message_type = _PAGE
_WIKI_PAGESENTRY.containing_type = _WIKI
_WIKI.fields_by_name['pages'].message_type = _WIKI_PAGESENTRY
_DATA_NAMEDENTITIESENTRY.fields_by_name['value'].message_type = _NAMEDENTITY
_DATA_NAMEDENTITIESENTRY.containing_type = _DATA
_DATA_TAGGEDPAGESENTRY.fields_by_name['value'].message_type = _TAGGEDNAMEDENTITIES
_DATA_TAGGEDPAGESENTRY.containing_type = _DATA
_DATA.fields_by_name['namedEntities'].message_type = _DATA_NAMEDENTITIESENTRY
_DATA.fields_by_name['taggedPages'].message_type = _DATA_TAGGEDPAGESENTRY
DESCRIPTOR.message_types_by_name['Wiki'] = _WIKI
DESCRIPTOR.message_types_by_name['Page'] = _PAGE
DESCRIPTOR.message_types_by_name['Data'] = _DATA
DESCRIPTOR.message_types_by_name['NamedEntity'] = _NAMEDENTITY
DESCRIPTOR.message_types_by_name['TaggedNamedEntities'] = _TAGGEDNAMEDENTITIES
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Wiki = _reflection.GeneratedProtocolMessageType('Wiki', (_message.Message,), dict(

  PagesEntry = _reflection.GeneratedProtocolMessageType('PagesEntry', (_message.Message,), dict(
    DESCRIPTOR = _WIKI_PAGESENTRY,
    __module__ = 'nn_auto_pb2'
    # @@protoc_insertion_point(class_scope:protobuf.Wiki.PagesEntry)
    ))
  ,
  DESCRIPTOR = _WIKI,
  __module__ = 'nn_auto_pb2'
  # @@protoc_insertion_point(class_scope:protobuf.Wiki)
  ))
_sym_db.RegisterMessage(Wiki)
_sym_db.RegisterMessage(Wiki.PagesEntry)

Page = _reflection.GeneratedProtocolMessageType('Page', (_message.Message,), dict(
  DESCRIPTOR = _PAGE,
  __module__ = 'nn_auto_pb2'
  # @@protoc_insertion_point(class_scope:protobuf.Page)
  ))
_sym_db.RegisterMessage(Page)

Data = _reflection.GeneratedProtocolMessageType('Data', (_message.Message,), dict(

  NamedEntitiesEntry = _reflection.GeneratedProtocolMessageType('NamedEntitiesEntry', (_message.Message,), dict(
    DESCRIPTOR = _DATA_NAMEDENTITIESENTRY,
    __module__ = 'nn_auto_pb2'
    # @@protoc_insertion_point(class_scope:protobuf.Data.NamedEntitiesEntry)
    ))
  ,

  TaggedPagesEntry = _reflection.GeneratedProtocolMessageType('TaggedPagesEntry', (_message.Message,), dict(
    DESCRIPTOR = _DATA_TAGGEDPAGESENTRY,
    __module__ = 'nn_auto_pb2'
    # @@protoc_insertion_point(class_scope:protobuf.Data.TaggedPagesEntry)
    ))
  ,
  DESCRIPTOR = _DATA,
  __module__ = 'nn_auto_pb2'
  # @@protoc_insertion_point(class_scope:protobuf.Data)
  ))
_sym_db.RegisterMessage(Data)
_sym_db.RegisterMessage(Data.NamedEntitiesEntry)
_sym_db.RegisterMessage(Data.TaggedPagesEntry)

NamedEntity = _reflection.GeneratedProtocolMessageType('NamedEntity', (_message.Message,), dict(
  DESCRIPTOR = _NAMEDENTITY,
  __module__ = 'nn_auto_pb2'
  # @@protoc_insertion_point(class_scope:protobuf.NamedEntity)
  ))
_sym_db.RegisterMessage(NamedEntity)

TaggedNamedEntities = _reflection.GeneratedProtocolMessageType('TaggedNamedEntities', (_message.Message,), dict(
  DESCRIPTOR = _TAGGEDNAMEDENTITIES,
  __module__ = 'nn_auto_pb2'
  # @@protoc_insertion_point(class_scope:protobuf.TaggedNamedEntities)
  ))
_sym_db.RegisterMessage(TaggedNamedEntities)


_WIKI_PAGESENTRY._options = None
_DATA_NAMEDENTITIESENTRY._options = None
_DATA_TAGGEDPAGESENTRY._options = None
# @@protoc_insertion_point(module_scope)
