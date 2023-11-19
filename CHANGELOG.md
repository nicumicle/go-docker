# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Allow to filter containers by name or IP
- Display more details when selecting a container ( status, IP, network, ports )

### Changed
- Code refactor
- Use JSON format for docker container ls and docker inspect
- Display only relevant actions when a container is selected

### Fixed
- Fix correct status when a container is restarting

## [0.0.1] - 2023-11-12

### Added
- Initial release