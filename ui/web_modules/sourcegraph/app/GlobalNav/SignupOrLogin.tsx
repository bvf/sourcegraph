import * as React from "react";

import {LocationStateToggleLink} from "sourcegraph/components/LocationStateToggleLink";

import * as base from "sourcegraph/components/styles/_base.css";
import * as AnalyticsConstants from "sourcegraph/util/constants/AnalyticsConstants";

export const SignupOrLogin = (props): JSX.Element => {

	const sx = Object.assign({},
		{
			textAlign: "center",
			display: "inline-block",
			lineHeight: "calc(100% - 1px)",
		},
		props.style
	);

	return(
		<div style={sx}>
			<LocationStateToggleLink href="/login" modalName="login" location={location}
				onToggle={(v) => v && AnalyticsConstants.Events.LoginModal_Initiated.logEvent({page_name: location.pathname, location_on_page: AnalyticsConstants.PAGE_LOCATION_GLOBAL_NAV})}>
				Log in
			</LocationStateToggleLink>
			<span className={base.mh1}> or </span>
			<LocationStateToggleLink href="/join" modalName="join" location={location}
				onToggle={(v) => v && AnalyticsConstants.Events.JoinModal_Initiated.logEvent({page_name: location.pathname, location_on_page: AnalyticsConstants.PAGE_LOCATION_GLOBAL_NAV})}>
				Sign up
			</LocationStateToggleLink>
		</div>
	);
};
