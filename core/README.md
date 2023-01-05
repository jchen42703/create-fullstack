# Create Fullstack Core <!-- omit in toc -->

Contains the main modules for UI/API/Fullstack template generation, configs, and template augmentation.

All of these modules are exposed to allow for plugin developers to easily develop their own templates and augmentations!

## Table of Contents <!-- omit in toc -->

- [`TemplatePipeline`](#templatepipeline)
- [`TemplateGenerator`](#templategenerator)
- [`TemplateAugmentor`](#templateaugmentor)

## `TemplatePipeline`

`TemplatePipeline`s are the main abstraction that the CLI uses to parse configs and use those configs for template generation and augmentation.

It composes various `TemplateGenerator`s and `TemplateAugmentor`s to generate and augment templates.

## `TemplateGenerator`

Responsible for downloading and filtering base templates.

## `TemplateAugmentor`

The shared augmentators are located in [`aug`](./aug/).

However, domain-specific (api/ui/fullstack) augmentations are located within their respective domains. For example, UI-specific tailwind augmentations are located in [`ui/tailwind.go`](ui/tailwind.go)
