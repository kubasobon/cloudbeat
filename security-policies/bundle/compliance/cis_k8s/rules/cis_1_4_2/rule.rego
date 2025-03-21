package compliance.cis_k8s.rules.cis_1_4_2

import data.compliance.policy.process.ensure_arguments_contain_key_value as audit
import future.keywords.if

finding := result if {
	audit.scheduler_filter
	result := audit.finding(audit.arg_contains("--bind-address", "127.0.0.1"))
}
