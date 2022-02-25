# Deployment Process

**User Story:** [EASI-1383](https://jiraent.cms.gov/browse/EASI-1383?filter=-1)

Evaluate the current deployment method.: 

- [ ] How application code is build
    - [ ] Application are build in a docker container deployed to ECS (AWS service) 
    - [ ] Databases are deployed 
- [ ] Backup and recovery procedures ( retention rate of data, backup intervals, rollback of application version) 
- [ ] Deployment method (This is a big one using GitHub workflows to deploy application with a lambda function to manage the ECS manager) 
- [ ] Test the current application Rollback method to review any pitfalls
*[decision drivers | forces]* <!-- optional -->

## Considered Alternatives

* Concourse
* GITHUB ACTIONS

## Decision Outcome

* Chosen Alternative: *[alternative 1]*
* *[justification.
  e.g., only alternative,
  which meets KO criterion decision driver
  | which resolves force force
  | ...
  | comes out best (see below)]*
* *[consequences. e.g.,
  negative impact on quality attribute,
  follow-up decisions required,
  ...]* <!-- optional -->

## Pros and Cons of the Alternatives <!-- optional -->

### Concourse

* `+` Pipeline-based CI that can do version control of files and applications 
* `+` Built in intergration with DevOps tools(Github, AWS, Vault)
* `+` 
* `-` Requires completely new pipeline builds and workflow built from strach 
* *[...]* <!-- numbers of pros and cons can vary -->

### *[alternative 2]*

* `+` *[argument 1 pro]*
* `+` *[argument 2 pro]*
* `-` *[argument 1 con]*
* *[...]* <!-- numbers of pros and cons can vary -->

### *[alternative 3]*

* `+` *[argument 1 pro]*
* `+` *[argument 2 pro]*
* `-` *[argument 1 con]*
* *[...]* <!-- numbers of pros and cons can vary -->
