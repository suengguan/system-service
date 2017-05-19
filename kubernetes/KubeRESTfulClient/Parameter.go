package KubeRESTfulClient

type Parameter struct {
	m_Namespace    string
	m_Name         string
	m_Options      map[string]string
	m_Fields       map[string]string
	m_Labels       map[string]string
	m_NotLabels    map[string]string
	m_InLabels     map[string][]string
	m_NotInLabels  map[string][]string
	m_Json         string
	m_ResourceType string
	m_SubPath      string
	m_IsVisitProxy bool
	m_IsSetWatcher bool
}

func (p *Parameter) BuildPath() string {
	var result string

	if p.m_IsVisitProxy {
		result += "/proxy"
	} else if p.m_IsSetWatcher {
		result += "/watch"
	}

	if len(p.m_Namespace) > 0 {
		result += "/namespaces/"
		result += p.m_Namespace
	}

	result += "/"
	result += p.m_ResourceType

	if len(p.m_Name) > 0 {
		result += "/"
		result += p.m_Name
	}

	if len(p.m_SubPath) > 0 {
		result += "/"
		result += p.m_SubPath
	}

	if len(p.m_Options) > 0 {
		result += "?"
		for k, v := range p.m_Options {
			result += k
			result += "="
			result += v
			result += "&&"
		}

	}

	if len(p.m_Labels) > 0 || len(p.m_NotLabels) > 0 || len(p.m_InLabels) > 0 || len(p.m_NotInLabels) > 0 || len(p.m_Fields) > 0 {
		labelSelectorStr := p.buildLabelSelector()
		fieldSelectorStr := p.buildFieldSelector()

		if (len(labelSelectorStr) + len(fieldSelectorStr)) > 0 {
			result += "?"
		}

		if len(labelSelectorStr) > 0 {
			result += "labelSelector="
			result += labelSelectorStr
			if len(fieldSelectorStr) > 0 {
				result += ","
			}
		}

		if len(fieldSelectorStr) > 0 {
			result += "fieldSelector="
			result += fieldSelectorStr
		}

	}

	return result
}

func (p *Parameter) buildLabelSelector() string {
	var result string

	if len(p.m_Labels) > 0 {
		for k, v := range p.m_Labels {
			if len(result) > 0 {
				result += ","
			}
			result += k
			result += "="
			result += v
		}
	}

	if len(p.m_NotLabels) > 0 {
		for k, v := range p.m_NotLabels {
			if len(result) > 0 {
				result += ","
			}
			result += k
			result += "!="
			result += v
		}
	}

	if len(p.m_InLabels) > 0 {
		for k, v := range p.m_InLabels {
			if len(result) > 0 {
				result += ","
			}
			result += k
			result += " in ("
			result += p.listToString(v, ",")
			result += ")"
		}
	}

	if len(p.m_NotInLabels) > 0 {
		for k, v := range p.m_NotInLabels {
			if len(result) > 0 {
				result += ","
			}
			result += k
			result += " notin ("
			result += p.listToString(v, ",")
			result += ")"
		}
	}

	return result
}

func (p *Parameter) buildFieldSelector() string {
	var result string

	if len(p.m_Fields) > 0 {
		for k, v := range p.m_NotLabels {
			if len(result) > 0 {
				result += ","
			}
			result += k
			result += "="
			result += v
		}
	}

	return result
}

func (p *Parameter) listToString(list []string, delim string) string {
	var result string

	for i := 0; i < len(list); i++ {
		if len(result) > 0 {
			result += delim
		}

		result += list[i]
	}

	return result
}

func (p *Parameter) SetNamespace(namespace string) {
	p.m_Namespace = namespace
}

func (p *Parameter) GetNamespace() string {
	return p.m_Namespace
}

func (p *Parameter) SetName(name string) {
	p.m_Name = name
}

func (p *Parameter) GetName() string {
	return p.m_Name
}

func (p *Parameter) SetOptions(options map[string]string) {
	p.m_Options = options
}

func (p *Parameter) GetOptions() map[string]string {
	return p.m_Options
}

func (p *Parameter) SetFields(fields map[string]string) {
	p.m_Fields = fields
}

func (p *Parameter) GetFields() map[string]string {
	return p.m_Fields
}

func (p *Parameter) SetLabels(labels map[string]string) {
	p.m_Labels = labels
}

func (p *Parameter) GetLabels() map[string]string {
	return p.m_Labels
}

func (p *Parameter) SetNotLabels(notLabels map[string]string) {
	p.m_NotLabels = notLabels
}

func (p *Parameter) GetNotLabels() map[string]string {
	return p.m_NotLabels
}

func (p *Parameter) SetInLabels(inLabels map[string][]string) {
	p.m_InLabels = inLabels
}

func (p *Parameter) GetInLabels() map[string][]string {
	return p.m_InLabels
}

func (p *Parameter) SetNotInLabels(notInLabels map[string][]string) {
	p.m_NotInLabels = notInLabels
}

func (p *Parameter) GetNotInLabels() map[string][]string {
	return p.m_NotInLabels
}

func (p *Parameter) SetJson(json string) {
	p.m_Json = json
}

func (p *Parameter) GetJson() string {
	return p.m_Json
}

func (p *Parameter) SetResourceType(resourceType string) {
	p.m_ResourceType = resourceType
}

func (p *Parameter) GetResourceType() string {
	return p.m_ResourceType
}

func (p *Parameter) SetSubPath(subPath string) {
	p.m_SubPath = subPath
}

func (p *Parameter) GetSubPath() string {
	return p.m_SubPath
}

func (p *Parameter) SetVisitProxy(isVisitProxy bool) {
	p.m_IsVisitProxy = isVisitProxy
}

func (p *Parameter) GetVisitProxy() bool {
	return p.m_IsVisitProxy
}

func (p *Parameter) SetSetWatcher(isSetWatcher bool) {
	p.m_IsSetWatcher = isSetWatcher
}

func (p *Parameter) GetSetWatcher() bool {
	return p.m_IsSetWatcher
}
