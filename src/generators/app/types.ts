export interface SpringMetadata {
    _links:       SpringMetadataLinks;
    dependencies: SpringMetadataDependencies;
    type:         Type;
    packaging:    BootVersion;
    javaVersion:  BootVersion;
    language:     BootVersion;
    bootVersion:  BootVersion;
    groupId:      ArtifactID;
    artifactId:   ArtifactID;
    version:      ArtifactID;
    name:         ArtifactID;
    description:  ArtifactID;
    packageName:  ArtifactID;
}

export interface SpringMetadataLinks {
    "gradle-project":        GradleBuildClass;
    "gradle-project-kotlin": GradleBuildClass;
    "gradle-build":          GradleBuildClass;
    "maven-project":         GradleBuildClass;
    "maven-build":           GradleBuildClass;
    dependencies:            GradleBuildClass;
}

export interface GradleBuildClass {
    href:      string;
    templated: boolean;
}

export interface ArtifactID {
    type:    string;
    default: string;
}

export interface BootVersion {
    type:    string;
    default: string;
    values:  BootVersionValue[];
}

export interface BootVersionValue {
    id:   string;
    name: string;
}

export interface SpringMetadataDependencies {
    type:   string;
    values: DependenciesValue[];
}

export interface DependenciesValue {
    name:   string;
    values: ValueValue[];
}

export interface ValueValue {
    id:            string;
    name:          string;
    description:   string;
    _links?:       ValueLinks;
    versionRange?: VersionRange;
}

export interface ValueLinks {
    reference?: Home[] | ReferenceClass;
    guide?:     Home[] | Home;
    home?:      Home;
    sample?:    Home;
}

export interface Home {
    href:   string;
    title?: string;
}

export interface ReferenceClass {
    href:       string;
    templated?: boolean;
    title?:     string;
}

export enum VersionRange {
    The340Release350M1 = "[3.4.0.RELEASE,3.5.0.M1)",
    The340Release360M1 = "[3.4.0.RELEASE,3.6.0.M1)",
}

export interface Type {
    type:    string;
    default: string;
    values:  TypeValue[];
}

export interface TypeValue {
    id:          string;
    name:        string;
    description: string;
    action:      string;
    tags:        Tags;
}

export interface Tags {
    build:    string;
    dialect?: string;
    format:   string;
}
