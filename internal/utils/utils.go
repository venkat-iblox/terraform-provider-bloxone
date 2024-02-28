package utils

import (
    "context"
    "fmt"

    datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
    "github.com/hashicorp/terraform-plugin-framework/diag"
    resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
    "github.com/hashicorp/terraform-plugin-log/tflog"
)

const ReadPageSizeLimit int32 = 1000

// Ptr is a helper routine that returns a pointer to given value.
func Ptr[T any](t T) *T {
    return &t
}

// DataSourceAttributeMap converts a map of resource schema attributes to data source schema attributes
func DataSourceAttributeMap(r map[string]resourceschema.Attribute, diags *diag.Diagnostics) map[string]datasourceschema.Attribute {
    d := map[string]datasourceschema.Attribute{}
    for k, v := range r {
        d[k] = DataSourceAttribute(k, v, diags)
    }
    return d
}

// DataSourceNestedAttributeObject converts a resource schema nested attribute object to data source schema nested attribute object
func DataSourceNestedAttributeObject(r resourceschema.NestedAttributeObject, diags *diag.Diagnostics) datasourceschema.NestedAttributeObject {
    return datasourceschema.NestedAttributeObject{
        Attributes: DataSourceAttributeMap(r.Attributes, diags),
        CustomType: r.CustomType,
        Validators: r.Validators,
    }
}

// DataSourceAttribute converts a resource schema attribute to data source schema attribute
func DataSourceAttribute(name string, val resourceschema.Attribute, diags *diag.Diagnostics) datasourceschema.Attribute {
    switch a := val.(type) {
    case resourceschema.BoolAttribute:
        return datasourceschema.BoolAttribute{
            CustomType:          a.CustomType,
            Required:            a.Required,
            Optional:            a.Optional,
            Computed:            a.Computed,
            Sensitive:           a.Sensitive,
            Description:         a.Description,
            MarkdownDescription: a.MarkdownDescription,
            DeprecationMessage:  a.DeprecationMessage,
            Validators:          a.Validators,
        }
    case resourceschema.StringAttribute:
        return datasourceschema.StringAttribute{
            CustomType:          a.CustomType,
            Required:            a.Required,
            Optional:            a.Optional,
            Computed:            a.Computed,
            Sensitive:           a.Sensitive,
            Description:         a.Description,
            MarkdownDescription: a.MarkdownDescription,
            DeprecationMessage:  a.DeprecationMessage,
            Validators:          a.Validators,
        }
    case resourceschema.Int64Attribute:
        return datasourceschema.Int64Attribute{
            CustomType:          a.CustomType,
            Required:            a.Required,
            Optional:            a.Optional,
            Computed:            a.Computed,
            Sensitive:           a.Sensitive,
            Description:         a.Description,
            MarkdownDescription: a.MarkdownDescription,
            DeprecationMessage:  a.DeprecationMessage,
            Validators:          a.Validators,
        }
    case resourceschema.Float64Attribute:
        return datasourceschema.Float64Attribute{
            CustomType:          a.CustomType,
            Required:            a.Required,
            Optional:            a.Optional,
            Computed:            a.Computed,
            Sensitive:           a.Sensitive,
            Description:         a.Description,
            MarkdownDescription: a.MarkdownDescription,
            DeprecationMessage:  a.DeprecationMessage,
            Validators:          a.Validators,
        }
    case resourceschema.NumberAttribute:
        return datasourceschema.NumberAttribute{
            CustomType:          a.CustomType,
            Required:            a.Required,
            Optional:            a.Optional,
            Computed:            a.Computed,
            Sensitive:           a.Sensitive,
            Description:         a.Description,
            MarkdownDescription: a.MarkdownDescription,
            DeprecationMessage:  a.DeprecationMessage,
            Validators:          a.Validators,
        }
    case resourceschema.ObjectAttribute:
        return datasourceschema.ObjectAttribute{
            AttributeTypes:      a.AttributeTypes,
            CustomType:          a.CustomType,
            Required:            a.Required,
            Optional:            a.Optional,
            Computed:            a.Computed,
            Sensitive:           a.Sensitive,
            Description:         a.Description,
            MarkdownDescription: a.MarkdownDescription,
            DeprecationMessage:  a.DeprecationMessage,
            Validators:          a.Validators,
        }
    case resourceschema.ListAttribute:
        return datasourceschema.ListAttribute{
            ElementType:         a.ElementType,
            CustomType:          a.CustomType,
            Required:            a.Required,
            Optional:            a.Optional,
            Computed:            a.Computed,
            Sensitive:           a.Sensitive,
            Description:         a.Description,
            MarkdownDescription: a.MarkdownDescription,
            DeprecationMessage:  a.DeprecationMessage,
            Validators:          a.Validators,
        }
    case resourceschema.ListNestedAttribute:
        return datasourceschema.ListNestedAttribute{
            NestedObject:        DataSourceNestedAttributeObject(a.NestedObject, diags),
            CustomType:          a.CustomType,
            Required:            a.Required,
            Optional:            a.Optional,
            Computed:            a.Computed,
            Sensitive:           a.Sensitive,
            Description:         a.Description,
            MarkdownDescription: a.MarkdownDescription,
            DeprecationMessage:  a.DeprecationMessage,
            Validators:          a.Validators,
        }
    case resourceschema.MapAttribute:
        return datasourceschema.MapAttribute{
            ElementType:         a.ElementType,
            CustomType:          a.CustomType,
            Required:            a.Required,
            Optional:            a.Optional,
            Computed:            a.Computed,
            Sensitive:           a.Sensitive,
            Description:         a.Description,
            MarkdownDescription: a.MarkdownDescription,
            DeprecationMessage:  a.DeprecationMessage,
            Validators:          a.Validators,
        }
    case resourceschema.MapNestedAttribute:
        return datasourceschema.MapNestedAttribute{
            NestedObject:        DataSourceNestedAttributeObject(a.NestedObject, diags),
            CustomType:          a.CustomType,
            Required:            a.Required,
            Optional:            a.Optional,
            Computed:            a.Computed,
            Sensitive:           a.Sensitive,
            Description:         a.Description,
            MarkdownDescription: a.MarkdownDescription,
            DeprecationMessage:  a.DeprecationMessage,
            Validators:          a.Validators,
        }
    case resourceschema.SetAttribute:
        return datasourceschema.SetAttribute{
            ElementType:         a.ElementType,
            CustomType:          a.CustomType,
            Required:            a.Required,
            Optional:            a.Optional,
            Computed:            a.Computed,
            Sensitive:           a.Sensitive,
            Description:         a.Description,
            MarkdownDescription: a.MarkdownDescription,
            DeprecationMessage:  a.DeprecationMessage,
            Validators:          a.Validators,
        }
    case resourceschema.SetNestedAttribute:
        return datasourceschema.SetNestedAttribute{
            NestedObject:        DataSourceNestedAttributeObject(a.NestedObject, diags),
            CustomType:          a.CustomType,
            Required:            a.Required,
            Optional:            a.Optional,
            Computed:            a.Computed,
            Sensitive:           a.Sensitive,
            Description:         a.Description,
            MarkdownDescription: a.MarkdownDescription,
            DeprecationMessage:  a.DeprecationMessage,
            Validators:          a.Validators,
        }
    case resourceschema.SingleNestedAttribute:
        return datasourceschema.SingleNestedAttribute{
            Attributes:          DataSourceAttributeMap(a.Attributes, diags),
            CustomType:          a.CustomType,
            Required:            a.Required,
            Optional:            a.Optional,
            Computed:            a.Computed,
            Sensitive:           a.Sensitive,
            Description:         a.Description,
            MarkdownDescription: a.MarkdownDescription,
            DeprecationMessage:  a.DeprecationMessage,
            Validators:          a.Validators,
        }
    }
    diags.AddError("Provider error",
        fmt.Sprintf("Failed to convert schema attribute of type '%T' for '%s'", val, name))
    return nil
}

func ReadWithPages[T any](read func(offset, limit int32) ([]T, error)) ([]T, error) {
    var allResults []T
    var offset int32 = 0

    for {
        results, err := read(offset, ReadPageSizeLimit)
        if err != nil {
            return nil, err
        }
        allResults = append(allResults, results...)
        if len(results) < int(ReadPageSizeLimit) {
            break
        }
        offset += ReadPageSizeLimit
    }

    return allResults, nil
}

// ToComputedAttributeMap converts a map of resource schema attributes to schema attributes with all fields set to "computed".
func ToComputedAttributeMap(r map[string]resourceschema.Attribute) map[string]resourceschema.Attribute {
    d := map[string]resourceschema.Attribute{}
    for k, v := range r {
        d[k] = ToComputedAttribute(k, v)
    }
    return d
}

// ToComputedNestedAttributeObject converts a resource schema nested attribute object to nested attribute object with all fields set to "computed".
func ToComputedNestedAttributeObject(r resourceschema.NestedAttributeObject) resourceschema.NestedAttributeObject {
    return resourceschema.NestedAttributeObject{
        Attributes: ToComputedAttributeMap(r.Attributes),
        CustomType: r.CustomType,
        Validators: r.Validators,
    }
}

// ToComputedAttribute converts a resource schema attribute having all attributes set to "computed".
func ToComputedAttribute(name string, val resourceschema.Attribute) resourceschema.Attribute {
    switch a := val.(type) {
    case resourceschema.StringAttribute:
        a.Required = false
        a.Optional = false
        a.Computed = true
        return a
    case resourceschema.BoolAttribute:
        a.Required = false
        a.Optional = false
        a.Computed = true
        return a
    case resourceschema.Int64Attribute:
        a.Required = false
        a.Optional = false
        a.Computed = true
        return a
    case resourceschema.Float64Attribute:
        a.Required = false
        a.Optional = false
        a.Computed = true
        return a
    case resourceschema.NumberAttribute:
        a.Required = false
        a.Optional = false
        a.Computed = true
        return a
    case resourceschema.ObjectAttribute:
        a.Required = false
        a.Optional = false
        a.Computed = true
        return a
    case resourceschema.ListAttribute:
        a.Required = false
        a.Optional = false
        a.Computed = true
        return a
    case resourceschema.ListNestedAttribute:
        a.NestedObject = ToComputedNestedAttributeObject(a.NestedObject)
        a.Required = false
        a.Optional = false
        a.Computed = true
        return a
    case resourceschema.MapAttribute:
        a.Required = false
        a.Optional = false
        a.Computed = true
        return a
    case resourceschema.MapNestedAttribute:
        a.NestedObject = ToComputedNestedAttributeObject(a.NestedObject)
        a.Required = false
        a.Optional = false
        a.Computed = true
        return a
    case resourceschema.SetAttribute:
        a.Required = false
        a.Optional = false
        a.Computed = true
        return a
    case resourceschema.SetNestedAttribute:
        a.NestedObject = ToComputedNestedAttributeObject(a.NestedObject)
        a.Required = false
        a.Optional = false
        a.Computed = true
        return a
    case resourceschema.SingleNestedAttribute:
        a.Attributes = ToComputedAttributeMap(a.Attributes)
        a.Required = false
        a.Optional = false
        a.Computed = true
        return a
    }

    tflog.Error(context.Background(), fmt.Sprintf("Failed to convert schema attribute of type '%T' for '%s'", val, name))
    return nil
}

// DefaultTagsHandler converts default_tags attribute to map[string]string
// supports only string values
func DefaultTagsHandler(m map[string]interface{}, d *diag.Diagnostics) map[string]string {
    newMap := make(map[string]string, len(m))
    for key, value := range m {
        strValue := ""

        switch v := value.(type) {
        case string:
            strValue = v
        default:
            d.AddError("Client error", fmt.Sprintf("Invalid type '%v' for default_tags %s value, only string supported", v, key))
            continue
        }

        newMap[key] = strValue
    }

    return newMap
}
