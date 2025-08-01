// #ddev-generated
package main

import (
    "fmt"
    "os"
    "strings"
    "github.com/goccy/go-yaml/ast"
    "github.com/goccy/go-yaml/parser"
    "github.com/goccy/go-yaml/token"
)

func main() {
    if err := run(); err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        os.Exit(1)
    }
}

func run() error {
    newHostname := strings.TrimSpace(os.Args[1])

    yamlFile := "config.yaml"
    data, err := os.ReadFile(yamlFile)
    if err != nil {
        return fmt.Errorf("reading file %s: %w", yamlFile, err)
    }

    file, err := parser.ParseBytes(data, parser.ParseComments)
    if err != nil {
        return fmt.Errorf("parsing YAML: %w", err)
    }

    if len(file.Docs) == 0 {
        return fmt.Errorf("no documents in YAML file")
    }

    root := file.Docs[0]
    if err := addHostname(root, newHostname); err != nil {
        return err
    }

    out := []byte(file.String())

    tempFile := yamlFile + ".tmp"
    if err := os.WriteFile(tempFile, out, 0644); err != nil {
        return fmt.Errorf("writing temporary file %s: %w", tempFile, err)
    }

    if err := os.Rename(tempFile, yamlFile); err != nil {
        os.Remove(tempFile)
        return fmt.Errorf("renaming temporary file: %w", err)
    }

    fmt.Printf("Successfully updated %s\n", yamlFile)
    return nil
}

func addHostname(node ast.Node, hostname string) error {
    doc, ok := node.(*ast.DocumentNode)
    if !ok {
        return fmt.Errorf("expected top-level YAML document")
    }

    if doc.Body == nil {
        return fmt.Errorf("document body is nil")
    }

    mapping, ok := doc.Body.(*ast.MappingNode)
    if !ok {
        return fmt.Errorf("expected top-level YAML mapping node, got %T", doc.Body)
    }

    var additionalHostnames *ast.SequenceNode

    for _, mappingValue := range mapping.Values {
        if mappingValue.Key == nil || mappingValue.Value == nil {
            continue
        }

        if keyNode, ok := mappingValue.Key.(*ast.StringNode); ok && keyNode.Value == "additional_hostnames" {
            if seqNode, ok := mappingValue.Value.(*ast.SequenceNode); ok {
                additionalHostnames = seqNode
                break
            }
        }
    }

    if additionalHostnames == nil {
        additionalHostnames = ast.Sequence(token.New("", "", &token.Position{}), false)
        newMappingValue := ast.MappingValue(
            token.New("", "", &token.Position{}),
            ast.String(token.New("additional_hostnames", "additional_hostnames", &token.Position{})),
            additionalHostnames,
        )

        mapping.Values = append(mapping.Values, newMappingValue)
    }

    for _, elem := range additionalHostnames.Values {
        if elem == nil {
            continue
        }
        if elemStr, ok := elem.(*ast.StringNode); ok && elemStr.Value == hostname {
            return nil
        }
    }

    newHostnameNode := ast.String(token.New(hostname, hostname, &token.Position{}))
    additionalHostnames.Values = append(additionalHostnames.Values, newHostnameNode)

    return nil
}
