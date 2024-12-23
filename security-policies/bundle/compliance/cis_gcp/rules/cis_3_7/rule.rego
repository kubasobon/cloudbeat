package compliance.cis_gcp.rules.cis_3_7

import data.compliance.lib.common
import data.compliance.policy.gcp.compute.ensure_fw_rule as audit
import data.compliance.policy.gcp.data_adapter
import future.keywords.if

finding := result if {
	# filter
	data_adapter.is_firewall_rule

	# set result
	result := common.generate_result_without_expected(
		common.calculate_result(is_rule_permissive),
		{"Firewall rule": input.resource},
	)
}

is_rule_permissive := audit.is_valid_fw_rule(3389) # RDP
