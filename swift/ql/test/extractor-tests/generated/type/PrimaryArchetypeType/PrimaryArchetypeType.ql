// generated by codegen/codegen.py
import codeql.swift.elements
import TestUtils

from
  PrimaryArchetypeType x, string getDiagnosticsName, Type getCanonicalType, string getName,
  Type getInterfaceType
where
  toBeTested(x) and
  not x.isUnknown() and
  getDiagnosticsName = x.getDiagnosticsName() and
  getCanonicalType = x.getCanonicalType() and
  getName = x.getName() and
  getInterfaceType = x.getInterfaceType()
select x, "getDiagnosticsName:", getDiagnosticsName, "getCanonicalType:", getCanonicalType,
  "getName:", getName, "getInterfaceType:", getInterfaceType