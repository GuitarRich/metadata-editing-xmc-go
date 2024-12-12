package components

import "guitarrich/xmcloud/sitecore/render/components"

func RegisterComponents() {
	components.RegisterComponent("Container", Container)
	components.RegisterComponent("CallToAction", CallToAction)
	components.RegisterComponent("Footer", Footer)
	components.RegisterComponent("Header", Header)
	components.RegisterComponent("Hero", Hero)
	components.RegisterComponent("Image", Image)
	components.RegisterComponent("Promo", Promo)
	components.RegisterComponent("RichText", RichText)
	components.RegisterComponent("Navigation", Navigation)
}
