package golang_excel

type ExcelDataSource struct {
  TemplateFileName string                 `json:"templateFileName"`
  Primary          map[string]interface{} `json:"primary"`
  ListSources      []ExcelListSource      `json:"listSources"`
}

func NewExcelDataSource(templateFileName string, primary map[string]interface{}, listSources []ExcelListSource) ExcelDataSource {
  return ExcelDataSource{
    TemplateFileName: templateFileName,
    Primary:          primary,
    ListSources:      listSources,
  }
}

type ExcelListSource struct {
  Name  string                   `json:"name"`
  Items []map[string]interface{} `json:"items"`
}

func NewExcelListSource(name string, items []map[string]interface{}) ExcelListSource {
  return ExcelListSource{
    Name:  name,
    Items: items,
  }
}

func NewExcelTableDataSource(name string, excelListSource ExcelListSource) ExcelDataSource {
  excelListSources := []ExcelListSource{excelListSource}
  return NewExcelDataSource(name, nil, excelListSources)
}

func NewExcelPrimaryDataSource(name string, primary map[string]interface{}) ExcelDataSource {
  return NewExcelDataSource(name, primary, []ExcelListSource{})
}
