package connector

import (
	v2 "github.com/conductorone/baton-sdk/pb/c1/connector/v2"
)

// The user resource type is for all user objects from the database.
var userResourceType = &v2.ResourceType{
	Id:          "user",
	DisplayName: "User",
	Traits:      []v2.ResourceType_Trait{v2.ResourceType_TRAIT_USER},
	Annotations: annotationsForUserResourceType(),
}

var accountResourceType = &v2.ResourceType{
	Id:          "account",
	DisplayName: "Account",
	Description: "Represents an Avalara account",
	Traits:      []v2.ResourceType_Trait{},
}

var permissionResourceType = &v2.ResourceType{
	Id:          "permission",
	DisplayName: "Permission",
	Description: "Represents an Avalara permission",
	Traits:      []v2.ResourceType_Trait{},
}

var roleResourceType = &v2.ResourceType{
	Id:          "role",
	DisplayName: "Role",
	Description: "Represents an Avalara security role",
	Traits:      []v2.ResourceType_Trait{v2.ResourceType_TRAIT_ROLE},
}
