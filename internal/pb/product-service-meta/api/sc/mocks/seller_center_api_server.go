// Code generated by mockery 2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	sc "gitlab.dg.ru/ocb/product-creation/product-storage-api/internal/pb/product-service-meta/api/sc"
)

// SellerCenterAPIServer is an autogenerated mock type for the SellerCenterAPIServer type
type SellerCenterAPIServer struct {
	mock.Mock
}

// CheckCategoryFinal provides a mock function with given fields: _a0, _a1
func (_m *SellerCenterAPIServer) CheckCategoryFinal(_a0 context.Context, _a1 *sc.CheckCategoryFinalRequest) (*sc.CheckCategoryFinalResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *sc.CheckCategoryFinalResponse
	if rf, ok := ret.Get(0).(func(context.Context, *sc.CheckCategoryFinalRequest) *sc.CheckCategoryFinalResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sc.CheckCategoryFinalResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sc.CheckCategoryFinalRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CheckSellerDescriptionCategoryExists provides a mock function with given fields: _a0, _a1
func (_m *SellerCenterAPIServer) CheckSellerDescriptionCategoryExists(_a0 context.Context, _a1 *sc.CheckSellerDescriptionCategoryExistsRequest) (*sc.CheckSellerDescriptionCategoryExistsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *sc.CheckSellerDescriptionCategoryExistsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *sc.CheckSellerDescriptionCategoryExistsRequest) *sc.CheckSellerDescriptionCategoryExistsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sc.CheckSellerDescriptionCategoryExistsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sc.CheckSellerDescriptionCategoryExistsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAttributesNamesMap provides a mock function with given fields: _a0, _a1
func (_m *SellerCenterAPIServer) GetAttributesNamesMap(_a0 context.Context, _a1 *sc.GetAttributesNamesMapRequest) (*sc.GetAttributesNamesMapResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *sc.GetAttributesNamesMapResponse
	if rf, ok := ret.Get(0).(func(context.Context, *sc.GetAttributesNamesMapRequest) *sc.GetAttributesNamesMapResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sc.GetAttributesNamesMapResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sc.GetAttributesNamesMapRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCategoriesAttributesWithHierarchy provides a mock function with given fields: _a0, _a1
func (_m *SellerCenterAPIServer) GetCategoriesAttributesWithHierarchy(_a0 context.Context, _a1 *sc.GetCategoriesAttributesWithHierarchyRequest) (*sc.GetCategoriesAttributesWithHierarchyResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *sc.GetCategoriesAttributesWithHierarchyResponse
	if rf, ok := ret.Get(0).(func(context.Context, *sc.GetCategoriesAttributesWithHierarchyRequest) *sc.GetCategoriesAttributesWithHierarchyResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sc.GetCategoriesAttributesWithHierarchyResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sc.GetCategoriesAttributesWithHierarchyRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCategoriesByIDs provides a mock function with given fields: _a0, _a1
func (_m *SellerCenterAPIServer) GetCategoriesByIDs(_a0 context.Context, _a1 *sc.GetCategoriesByIDsRequest) (*sc.GetCategoriesByIDsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *sc.GetCategoriesByIDsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *sc.GetCategoriesByIDsRequest) *sc.GetCategoriesByIDsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sc.GetCategoriesByIDsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sc.GetCategoriesByIDsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCategoryAttributesWithDictionaryValue provides a mock function with given fields: _a0, _a1
func (_m *SellerCenterAPIServer) GetCategoryAttributesWithDictionaryValue(_a0 context.Context, _a1 *sc.GetCategoryAttributesRequest) (*sc.GetCategoryAttributesResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *sc.GetCategoryAttributesResponse
	if rf, ok := ret.Get(0).(func(context.Context, *sc.GetCategoryAttributesRequest) *sc.GetCategoryAttributesResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sc.GetCategoryAttributesResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sc.GetCategoryAttributesRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCategoryBranchByID provides a mock function with given fields: _a0, _a1
func (_m *SellerCenterAPIServer) GetCategoryBranchByID(_a0 context.Context, _a1 *sc.GetCategoryBranchByIDRequest) (*sc.GetCategoryBranchByIDResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *sc.GetCategoryBranchByIDResponse
	if rf, ok := ret.Get(0).(func(context.Context, *sc.GetCategoryBranchByIDRequest) *sc.GetCategoryBranchByIDResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sc.GetCategoryBranchByIDResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sc.GetCategoryBranchByIDRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCategoryByTypeIDs provides a mock function with given fields: _a0, _a1
func (_m *SellerCenterAPIServer) GetCategoryByTypeIDs(_a0 context.Context, _a1 *sc.GetCategoryByTypeIDsRequest) (*sc.GetCategoryByTypeIDsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *sc.GetCategoryByTypeIDsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *sc.GetCategoryByTypeIDsRequest) *sc.GetCategoryByTypeIDsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sc.GetCategoryByTypeIDsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sc.GetCategoryByTypeIDsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCategoryTypeMapByCategoryIDs provides a mock function with given fields: _a0, _a1
func (_m *SellerCenterAPIServer) GetCategoryTypeMapByCategoryIDs(_a0 context.Context, _a1 *sc.GetCategoryTypeMapByCategoryIDsRequest) (*sc.GetCategoryTypeMapByCategoryIDsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *sc.GetCategoryTypeMapByCategoryIDsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *sc.GetCategoryTypeMapByCategoryIDsRequest) *sc.GetCategoryTypeMapByCategoryIDsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sc.GetCategoryTypeMapByCategoryIDsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sc.GetCategoryTypeMapByCategoryIDsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCategoryTypeMapByCategoryTypeIDs provides a mock function with given fields: _a0, _a1
func (_m *SellerCenterAPIServer) GetCategoryTypeMapByCategoryTypeIDs(_a0 context.Context, _a1 *sc.GetCategoryTypeMapByCategoryTypeIDsRequest) (*sc.GetCategoryTypeMapByCategoryTypeIDsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *sc.GetCategoryTypeMapByCategoryTypeIDsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *sc.GetCategoryTypeMapByCategoryTypeIDsRequest) *sc.GetCategoryTypeMapByCategoryTypeIDsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sc.GetCategoryTypeMapByCategoryTypeIDsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sc.GetCategoryTypeMapByCategoryTypeIDsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCommercialCategoriesByIDs provides a mock function with given fields: _a0, _a1
func (_m *SellerCenterAPIServer) GetCommercialCategoriesByIDs(_a0 context.Context, _a1 *sc.GetCommercialCategoriesByIDsRequest) (*sc.GetCommercialCategoriesByIDsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *sc.GetCommercialCategoriesByIDsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *sc.GetCommercialCategoriesByIDsRequest) *sc.GetCommercialCategoriesByIDsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sc.GetCommercialCategoriesByIDsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sc.GetCommercialCategoriesByIDsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCommercialCategoriesByMetazonIDs provides a mock function with given fields: _a0, _a1
func (_m *SellerCenterAPIServer) GetCommercialCategoriesByMetazonIDs(_a0 context.Context, _a1 *sc.GetCommercialCategoriesByMetazonIDsRequest) (*sc.GetCommercialCategoriesByMetazonIDsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *sc.GetCommercialCategoriesByMetazonIDsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *sc.GetCommercialCategoriesByMetazonIDsRequest) *sc.GetCommercialCategoriesByMetazonIDsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sc.GetCommercialCategoriesByMetazonIDsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sc.GetCommercialCategoriesByMetazonIDsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCommercialCategoriesByTypeIDs provides a mock function with given fields: _a0, _a1
func (_m *SellerCenterAPIServer) GetCommercialCategoriesByTypeIDs(_a0 context.Context, _a1 *sc.GetCommercialCategoriesByTypeIDsRequest) (*sc.GetCommercialCategoriesByTypeIDsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *sc.GetCommercialCategoriesByTypeIDsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *sc.GetCommercialCategoriesByTypeIDsRequest) *sc.GetCommercialCategoriesByTypeIDsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sc.GetCommercialCategoriesByTypeIDsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sc.GetCommercialCategoriesByTypeIDsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCommercialCategoriesMap provides a mock function with given fields: _a0, _a1
func (_m *SellerCenterAPIServer) GetCommercialCategoriesMap(_a0 context.Context, _a1 *sc.GetCommercialCategoriesMapRequest) (*sc.GetCommercialCategoriesMapResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *sc.GetCommercialCategoriesMapResponse
	if rf, ok := ret.Get(0).(func(context.Context, *sc.GetCommercialCategoriesMapRequest) *sc.GetCommercialCategoriesMapResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sc.GetCommercialCategoriesMapResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sc.GetCommercialCategoriesMapRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCommercialCategoriesTree provides a mock function with given fields: _a0, _a1
func (_m *SellerCenterAPIServer) GetCommercialCategoriesTree(_a0 context.Context, _a1 *sc.GetCommercialCategoriesTreeRequest) (*sc.GetCommercialCategoriesTreeResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *sc.GetCommercialCategoriesTreeResponse
	if rf, ok := ret.Get(0).(func(context.Context, *sc.GetCommercialCategoriesTreeRequest) *sc.GetCommercialCategoriesTreeResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sc.GetCommercialCategoriesTreeResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sc.GetCommercialCategoriesTreeRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDescriptionCategoryDeep provides a mock function with given fields: _a0, _a1
func (_m *SellerCenterAPIServer) GetDescriptionCategoryDeep(_a0 context.Context, _a1 *sc.GetDescriptionCategoryDeepRequest) (*sc.GetDescriptionCategoryDeepResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *sc.GetDescriptionCategoryDeepResponse
	if rf, ok := ret.Get(0).(func(context.Context, *sc.GetDescriptionCategoryDeepRequest) *sc.GetDescriptionCategoryDeepResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sc.GetDescriptionCategoryDeepResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sc.GetDescriptionCategoryDeepRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDictionaryBrands provides a mock function with given fields: _a0, _a1
func (_m *SellerCenterAPIServer) GetDictionaryBrands(_a0 context.Context, _a1 *sc.GetDictionaryBrandsRequest) (*sc.GetDictionaryBrandsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *sc.GetDictionaryBrandsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *sc.GetDictionaryBrandsRequest) *sc.GetDictionaryBrandsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sc.GetDictionaryBrandsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sc.GetDictionaryBrandsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDictionaryByIDs provides a mock function with given fields: _a0, _a1
func (_m *SellerCenterAPIServer) GetDictionaryByIDs(_a0 context.Context, _a1 *sc.GetDictionaryByIDsRequest) (*sc.GetDictionaryByIDsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *sc.GetDictionaryByIDsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *sc.GetDictionaryByIDsRequest) *sc.GetDictionaryByIDsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sc.GetDictionaryByIDsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sc.GetDictionaryByIDsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDictionaryValueBatch provides a mock function with given fields: _a0, _a1
func (_m *SellerCenterAPIServer) GetDictionaryValueBatch(_a0 context.Context, _a1 *sc.GetDictionaryValueBatchRequest) (*sc.GetDictionaryValueBatchResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *sc.GetDictionaryValueBatchResponse
	if rf, ok := ret.Get(0).(func(context.Context, *sc.GetDictionaryValueBatchRequest) *sc.GetDictionaryValueBatchResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sc.GetDictionaryValueBatchResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sc.GetDictionaryValueBatchRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDictionaryValuesByAttributeIDAndRsIDList provides a mock function with given fields: _a0, _a1
func (_m *SellerCenterAPIServer) GetDictionaryValuesByAttributeIDAndRsIDList(_a0 context.Context, _a1 *sc.GetDictionaryValuesByAttributeIDAndRsIDListRequest) (*sc.GetDictionaryValuesByAttributeIDAndRsIDListResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *sc.GetDictionaryValuesByAttributeIDAndRsIDListResponse
	if rf, ok := ret.Get(0).(func(context.Context, *sc.GetDictionaryValuesByAttributeIDAndRsIDListRequest) *sc.GetDictionaryValuesByAttributeIDAndRsIDListResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sc.GetDictionaryValuesByAttributeIDAndRsIDListResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sc.GetDictionaryValuesByAttributeIDAndRsIDListRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDictionaryValuesByDescCategoryAttribute provides a mock function with given fields: _a0, _a1
func (_m *SellerCenterAPIServer) GetDictionaryValuesByDescCategoryAttribute(_a0 context.Context, _a1 *sc.GetDictionaryValuesByDescCategoryAttributeRequest) (*sc.GetDictionaryValuesByDescCategoryAttributeResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *sc.GetDictionaryValuesByDescCategoryAttributeResponse
	if rf, ok := ret.Get(0).(func(context.Context, *sc.GetDictionaryValuesByDescCategoryAttributeRequest) *sc.GetDictionaryValuesByDescCategoryAttributeResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sc.GetDictionaryValuesByDescCategoryAttributeResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sc.GetDictionaryValuesByDescCategoryAttributeRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDictionaryValuesByDictionaryExternalIDs provides a mock function with given fields: _a0, _a1
func (_m *SellerCenterAPIServer) GetDictionaryValuesByDictionaryExternalIDs(_a0 context.Context, _a1 *sc.GetDictionaryValuesByDictionaryExternalIDsRequest) (*sc.GetDictionaryValuesByDictionaryExternalIDsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *sc.GetDictionaryValuesByDictionaryExternalIDsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *sc.GetDictionaryValuesByDictionaryExternalIDsRequest) *sc.GetDictionaryValuesByDictionaryExternalIDsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sc.GetDictionaryValuesByDictionaryExternalIDsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sc.GetDictionaryValuesByDictionaryExternalIDsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDictionaryValuesByDictionaryKeyAndRsIDList provides a mock function with given fields: _a0, _a1
func (_m *SellerCenterAPIServer) GetDictionaryValuesByDictionaryKeyAndRsIDList(_a0 context.Context, _a1 *sc.GetDictionaryValuesByDictionaryKeyAndRsIDListRequest) (*sc.GetDictionaryValuesByDictionaryKeyAndRsIDListResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *sc.GetDictionaryValuesByDictionaryKeyAndRsIDListResponse
	if rf, ok := ret.Get(0).(func(context.Context, *sc.GetDictionaryValuesByDictionaryKeyAndRsIDListRequest) *sc.GetDictionaryValuesByDictionaryKeyAndRsIDListResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sc.GetDictionaryValuesByDictionaryKeyAndRsIDListResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sc.GetDictionaryValuesByDictionaryKeyAndRsIDListRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDictionaryValuesByIDs provides a mock function with given fields: _a0, _a1
func (_m *SellerCenterAPIServer) GetDictionaryValuesByIDs(_a0 context.Context, _a1 *sc.GetDictionaryValuesByIDsRequest) (*sc.GetDictionaryValuesByIDsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *sc.GetDictionaryValuesByIDsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *sc.GetDictionaryValuesByIDsRequest) *sc.GetDictionaryValuesByIDsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sc.GetDictionaryValuesByIDsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sc.GetDictionaryValuesByIDsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFinalUnmappedDescriptionCategories provides a mock function with given fields: _a0, _a1
func (_m *SellerCenterAPIServer) GetFinalUnmappedDescriptionCategories(_a0 context.Context, _a1 *sc.GetFinalUnmappedDescriptionCategoriesRequest) (*sc.GetFinalUnmappedDescriptionCategoriesResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *sc.GetFinalUnmappedDescriptionCategoriesResponse
	if rf, ok := ret.Get(0).(func(context.Context, *sc.GetFinalUnmappedDescriptionCategoriesRequest) *sc.GetFinalUnmappedDescriptionCategoriesResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sc.GetFinalUnmappedDescriptionCategoriesResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sc.GetFinalUnmappedDescriptionCategoriesRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMappedCommercialCategoriesByDescriptionIDs provides a mock function with given fields: _a0, _a1
func (_m *SellerCenterAPIServer) GetMappedCommercialCategoriesByDescriptionIDs(_a0 context.Context, _a1 *sc.GetMappedCommercialCategoriesByDescriptionIDsRequest) (*sc.GetMappedCommercialCategoriesByDescriptionIDsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *sc.GetMappedCommercialCategoriesByDescriptionIDsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *sc.GetMappedCommercialCategoriesByDescriptionIDsRequest) *sc.GetMappedCommercialCategoriesByDescriptionIDsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sc.GetMappedCommercialCategoriesByDescriptionIDsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sc.GetMappedCommercialCategoriesByDescriptionIDsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSellerCategoryLevelsByCategoryIDs provides a mock function with given fields: _a0, _a1
func (_m *SellerCenterAPIServer) GetSellerCategoryLevelsByCategoryIDs(_a0 context.Context, _a1 *sc.GetSellerCategoryLevelsByCategoryIDsRequest) (*sc.GetSellerCategoryLevelsByCategoryIDsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *sc.GetSellerCategoryLevelsByCategoryIDsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *sc.GetSellerCategoryLevelsByCategoryIDsRequest) *sc.GetSellerCategoryLevelsByCategoryIDsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sc.GetSellerCategoryLevelsByCategoryIDsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sc.GetSellerCategoryLevelsByCategoryIDsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSellerCategoryTree provides a mock function with given fields: _a0, _a1
func (_m *SellerCenterAPIServer) GetSellerCategoryTree(_a0 context.Context, _a1 *sc.GetSellerCategoryTreeRequest) (*sc.GetSellerCategoryTreeResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *sc.GetSellerCategoryTreeResponse
	if rf, ok := ret.Get(0).(func(context.Context, *sc.GetSellerCategoryTreeRequest) *sc.GetSellerCategoryTreeResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sc.GetSellerCategoryTreeResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sc.GetSellerCategoryTreeRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSellerDescriptionCategoriesByIDs provides a mock function with given fields: _a0, _a1
func (_m *SellerCenterAPIServer) GetSellerDescriptionCategoriesByIDs(_a0 context.Context, _a1 *sc.GetSellerDescriptionCategoriesByIDsRequest) (*sc.GetSellerDescriptionCategoriesByIDsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *sc.GetSellerDescriptionCategoriesByIDsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *sc.GetSellerDescriptionCategoriesByIDsRequest) *sc.GetSellerDescriptionCategoriesByIDsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sc.GetSellerDescriptionCategoriesByIDsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sc.GetSellerDescriptionCategoriesByIDsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSellerDescriptionCategoriesFinalByDescriptionID provides a mock function with given fields: _a0, _a1
func (_m *SellerCenterAPIServer) GetSellerDescriptionCategoriesFinalByDescriptionID(_a0 context.Context, _a1 *sc.GetSellerDescriptionCategoriesFinalByDescriptionIDRequest) (*sc.GetSellerDescriptionCategoriesFinalByDescriptionIDResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *sc.GetSellerDescriptionCategoriesFinalByDescriptionIDResponse
	if rf, ok := ret.Get(0).(func(context.Context, *sc.GetSellerDescriptionCategoriesFinalByDescriptionIDRequest) *sc.GetSellerDescriptionCategoriesFinalByDescriptionIDResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sc.GetSellerDescriptionCategoriesFinalByDescriptionIDResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sc.GetSellerDescriptionCategoriesFinalByDescriptionIDRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSellerDescriptionCategoryBranchByID provides a mock function with given fields: _a0, _a1
func (_m *SellerCenterAPIServer) GetSellerDescriptionCategoryBranchByID(_a0 context.Context, _a1 *sc.GetSellerDescriptionCategoryBranchByIDRequest) (*sc.GetSellerDescriptionCategoryBranchByIDResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *sc.GetSellerDescriptionCategoryBranchByIDResponse
	if rf, ok := ret.Get(0).(func(context.Context, *sc.GetSellerDescriptionCategoryBranchByIDRequest) *sc.GetSellerDescriptionCategoryBranchByIDResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sc.GetSellerDescriptionCategoryBranchByIDResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sc.GetSellerDescriptionCategoryBranchByIDRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchBrandByName provides a mock function with given fields: _a0, _a1
func (_m *SellerCenterAPIServer) SearchBrandByName(_a0 context.Context, _a1 *sc.SearchBrandByNameRequest) (*sc.SearchBrandByNameResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *sc.SearchBrandByNameResponse
	if rf, ok := ret.Get(0).(func(context.Context, *sc.SearchBrandByNameRequest) *sc.SearchBrandByNameResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sc.SearchBrandByNameResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sc.SearchBrandByNameRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchCommercialCategoriesByName provides a mock function with given fields: _a0, _a1
func (_m *SellerCenterAPIServer) SearchCommercialCategoriesByName(_a0 context.Context, _a1 *sc.SearchCommercialCategoriesByNameRequest) (*sc.SearchCommercialCategoriesByNameResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *sc.SearchCommercialCategoriesByNameResponse
	if rf, ok := ret.Get(0).(func(context.Context, *sc.SearchCommercialCategoriesByNameRequest) *sc.SearchCommercialCategoriesByNameResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sc.SearchCommercialCategoriesByNameResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sc.SearchCommercialCategoriesByNameRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ToDescriptionCategory provides a mock function with given fields: _a0, _a1
func (_m *SellerCenterAPIServer) ToDescriptionCategory(_a0 context.Context, _a1 *sc.ToDescriptionCategoryRequest) (*sc.ToDescriptionCategoryResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *sc.ToDescriptionCategoryResponse
	if rf, ok := ret.Get(0).(func(context.Context, *sc.ToDescriptionCategoryRequest) *sc.ToDescriptionCategoryResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sc.ToDescriptionCategoryResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sc.ToDescriptionCategoryRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mustEmbedUnimplementedSellerCenterAPIServer provides a mock function with given fields:
func (_m *SellerCenterAPIServer) mustEmbedUnimplementedSellerCenterAPIServer() {
	_m.Called()
}