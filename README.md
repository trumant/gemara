# SCI: Simplified Compliance Infrastructure

The _Simplified Compliance Infrastructure_ provides a model to describe the categories of compliance activities, how they interact, and the schemas to allow for interoperability between them.

In order to better facilitate cross-functional communication, the SCI Model seeks to outline the categorical layers of activities related to automated governance.

We will begin by establishing the overall model, and then the following sections will contain detailed breakdowns of each categorical layer, with examples.

This document assumes that the reader is trained in governance, risk, compliance, or cybersecurity— and therefore understands or can find definitions of concepts or terms that are infrequently used herein.

For the purpose of this document, "organization" may refer to a business or an organizational unit within it.

## The Model

Each layer in the model builds upon the lower layer, though in higher-level use cases you may find examples where multiple lower layers are brought into a higher level together. Examples and clarifications can be found in the respective sections below.

| Layer | Name | Description |
|-------|------|-------------|
| 1 | Rules | High-level guidance on cybersecurity measures |
| 2 | Controls | Technology-specific, threat-informed security controls |
| 3 | Policy | Risk-informed governance rules tailored to an organization |
| 4 | Evaluation | Inspection of code, configurations, and deployments |
| 5 | Enforcement | Prevention or remediation based on assessment findings |
| 6 | Audit | Monitoring, logging, and reporting on deployed resource state |

### Layer 1: Rules

The Rules layer is the lowest level of the SCI Model. Activities in this layer provide high-level guidance on cybersecurity measures. Rules are typically developed by industry groups, government agencies, or international standards bodies. They are intended to be used as a starting point for organizations to develop their own cybersecurity programs.

Rules frameworks or standards occasionally express their guidance using the term "controls" — these should be understood as Layer 1 Controls in the event that the term appears to conflict with Layer 2. Rules are high-level, abstract controls that should be referenced by technology-specific Layer 2 Controls.

Examples include the NIST Cybersecurity Framework, ISO 27001, PCI DSS, HIPPA, GDPR, and CRA — among others.

### Layer 2: Controls

Activities in the Control layer produce technology-specific, threat-informed security controls. Controls are the specific guardrails that organizations put in place to protect their information systems. They are typically informed by the best practices and industry standards which are produced in Layer 1.

Layer 2 Controls are typically developed by an organization for its own purposes, or for general use by industry groups, government agencies, or international standards bodies. These should be combined with organizational risk considerations to form Layer 3 controls.

The recommended process for developing Layer 2 Controls is to first assess the technology's capabilities, then identify threats to those capabilities, and finally develop controls to mitigate those threats.

Examples include CIS Benchmarks, FINOS Common Cloud Controls, and the OSPS Baseline.

#### Layer 2 Schema

The SCI [Layer 2 Schema](./schemas/layer-2.cue) provides methods for expressing Layer 2 controls in a machine-readable format. 

The schema allows controls to be mapped to threats or Layer 1 controls by their unique identifiers. Threats may also be expressed in the schema, with mappings to the technology-specific capabilities which may be vulnerable to the threat.

The SCI go module provides Layer 2 support for ingesting documents that follow this schema.

### Layer 3: Policy

Activities in the Policy layer provide risk-informed governance rules tailored to an organization. Policies are the rules that organizations put in place to govern their information systems. They are typically informed by risks or based on best practices and industry standards.

Layer 3 controls are typically developed by an organization for its own purposes, to compile into organizational policies. Policies cannot be properly developed without consideration for organization-specific risk appetite and risk-acceptance. These should be used as a starting point for Layer 4 assessments.

### Layer 4: Evaluation

Activities in the Evaluation layer provide inspection of code, configurations, and deployments. Those elements are part of the _software development lifecycle_ which is not represented in this model.

Evaluation activities may be built based on outputs from Layers 2 or 3. While automated assessments are often developed by vendors or industry groups, proper evaluation should also be able to ingest organizational policies to custom-tailor the assessment to the needs of the compliance program.

#### Layer 4 Schema

The SCI [Layer 4 Schema)(./schemas/layer-4.cue) provides methods for expressing Layer 4 evaluation results in a machine readable format.

The schema allows evaluations to be mapped to Layer 2 controls by their unique identifiers.

The SCI go module provides Layer 2 support for writing assessments that can write results in this schema.

### Layer 5: Enforcement

Activities in the Enforcement layer provide prevention or remediation based on assessment findings.

Enforcement may be built based on Layer 4 assessments. This layer ensures that the organization is able to take action based on the results of the assessments, such as by blocking the deployment of a resource that does not meet the organization's policies.

### Layer 6: Audit

Activities in the Audit layer provide monitoring, logging, and reporting on deployed resource state.

Auditing may be built based on Layer 4 assessment results, informed by Layer 5 enforcement remediation activities, while considering other requirements of the organization — such as cost monitoring. This layer ensures that the organization is able to monitor the state of deployed resources and report on their compliance with the organization's policies.

## Tooling

The SCI Model is intended to be used in conjunction with tooling that can help automate the compliance process.

This section will be updated as SCI-compatible tooling becomes available for production use cases.
