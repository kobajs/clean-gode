# clean-gode

## Table of Contents

1. [Introduction](#introduction)
2. [Meaningful Names](#meaningful-names)
3. [Functions](#functions)
4. [Objects and Data Structures](#objects-and-data-structures)
5. [Classes](#classes)
6. [SOLID](#solid)
7. [Testing](#testing)
8. [Concurrency](#concurrency)
9. [Error Handling](#error-handling)
10. [Formatting](#formatting)
11. [Comments](#comments)
12. [Translation](#translation)

## Introduction

![Humorous image of software quality estimation as a count of how many expletives
you shout when reading code](https://www.osnews.com/images/comics/wtfm.jpg)

Software engineering principles, from Robert C.Martin's book [_Clean Code_](https://www.amazon.com/Clean-Code-Handbook-Software-Craftsmanship-ebook/dp/B001GSTOAM), adapted for [Golang](https://golang.org/). Here you will find useful DO's and DONT's that focus on good quality of code, producing [readable, reusable and refactorable](https://github.com/ryanmcdermott/3rs-of-software-architecture) software in Golang.

First of all, it could have some disagreement about the recommendations below and it's fine. Even the author says that his book is not a list of strict rules, but it's extremally important to notice that his creation is from good years of experience on area, discussions with high level software engineers and tons of research.

Personally, what I like from this book is the way that tells to programmers that coding is just like poetry, it needs to be pleasant to read, organized harmonically and give a story to the user. It's a care with the expressiveness of the author.

Also, it does not guarantee that after the read, you will be the best developer. Instead of that, practicing and collaborating with your own experiences through this rules (and others) for years, you will became a better developer.

## Meaningful Names

The name of a variable, function, or class, should answer all the big questions. It should tell why it exists, what it does and how it is used.

### Use Intention-Revealing Names

If a name requires a comment, then the name does not reveal its intent.

**Bad:**

```go
var cs string // Company Street
```

**Good:**

```go
var companyAddress string
```

### Avoid Disinformation

**Avoid names with structure at suffix if it is not one: accountList (need to be a List or should be called only accounts).**
Avoid names like "O"/0 and "l"/1. Font can trick the reader.
Avoid poor names used in systems: hp could be Unix system of hypotenuse.

### Make Meaningful Distinctions

It's hard to decide which to use if it's hard to distinct between similar or synonym meaning.

**Bad:**

```go
var product Product
var productData Product
var productInfo Product

var user User
var theUser User
var aUser User

var order1 Order
var order2 Order

var valueInt int
var nameString string

func getDisabledUser() {}
func getDisabledUsers() {}
func getDisabledUserList() {}
```

**Good:**

```go
var productOrigin Address
var productDestiny Address
```

### Use Pronounceable Names

When a name is pronounceable, it's easier to associate to something, it turns to be closer of real world understanding

**Bad:**

```go
var genymdhms time.Time
var uaid string
```

**Good:**

```go
var generationTimestamp time.Time
var userAccountId string
```

### Use Searchable Names

Single-letter names and numeric constants are harder to locate them. Try to assign with a meaning variable to be easier to found

**Bad:**

```go
const n int = 5
const h int = 8
```

**Good:**

```go
const WORK_DAYS_PER_WEEK int = 5
const WORK_HOURS int = 8
```

### Avoid Encondings / Hungarian Notations / Member prefixes

Encoding means that it's another language to be learned and simply burns unnecesarily our brain.

There's a dark time that hungarian notation was useful. With the modern IDE's, it's not necessary anymore.

Also member prefixes have high changes to be ignored.

**Bad:**

```go
const c_maxMoves int = 20

var strStreet string
var i64CoordX int64

var u_user User
```

**Good:**

```go
const characterMaxMoves int = 20

var street string
var coordX int64

var user User
```

### Interfaces and Implementations

Interfaces allow you to define behavior without exposing internal implementation.

Those that designed this incredible language likes the quote "if it walks like a duck, swims like a duck and quacks like a duck, then it’s a duck".
Note: Go delivers interfaces automagically and implicitly.

If we have interfaces, we also have implementations by a concrete class and there is special cases that their names are similars. So, here we could have a bit of encoding. Saying that, it's better to deliver a pleasant interface name than the implementation.

**Bad:**

```go
type IShapeFactory interface { }
type shapeFactory struct { }
```

**Good:**

```go
type ShapeFactory interface { }
type ShapeFactoryImpl struct { }
```

### Avoid Mental Mapping

Readers shouldn't have to mentally translate names into other names they already know. E.g: Short variable names like _a_, _j_ and _i_ are useful when the function is shorten, otherwise became an unnecesary mental mapping.

### Class and Object Names

Should have noun or noun phrases names like _Chapter_, _Process_, _Step_, and _AddressParser_. Avoid words like _Manager_, _Processor_, _Data_, or _Info_ and verbs.

### Method Names

Should have verb or verb phrase names like _addStep_, _deletePage_, or _create_. Accessors, mutators, and predicates should be named for their value and prefixed with _get_, _set_, and _is_.

### Don't Be Cute

Choose clarity over entertainment

**Bad:**

```go
func whack() { }
func eatMyShort() { }
func goodbyeUser() { }
```

**Good:**

```go
func kill() { }
func abort() { }
func removeUser() { }
```

### Pick One Word per Concept

Avoid different method names with _fetch_, _retrieve_, and _get_ as prefix. Get one of them and stick with it.
A consistent lexicon is a great boon to the programmers who must use your code.

**Bad:**

```go
func (a Address) getAddress(id int) { }
func (p Person) retrievePerson(id int) { }
func (c Company) fetchCompany(id int) { }
```

**Good:**

```go
func (a Address) getAddress(id int) { }
func (p Person) getPerson(id int) { }
func (c Company) getCompany(id int) { }
```

### Don't Pun

Avoid using the same word for two purposes.
Add to create or concatenate should be divided into insert and append.

**Bad:**

```go
func (a Address) addAddress(a Address) {
  db.Create(&a)
}

func (pl *PersonList) addPerson(p Person) {
  *pl = append(*pl, p)
}
```

**Good:**

```go
func (a Address) insertAddress(a Address) {
  db.Create(&a)
}

func (pl *PersonList) appendPerson(p Person) {
  *pl = append(*pl, p)
}
```

### Use Solution Domain Names

Avoid using the same word for two purposes.
Add to create or concatenate should be divided into insert and append.

**Bad:**

```go
func (a Address) addAddress(a Address) {
  db.Create(&a)
}

func (pl *PersonList) addPerson(p Person) {
  *pl = append(*pl, p)
}
```

**Good:**

```go
func (a Address) insertAddress(a Address) {
  db.Create(&a)
}

func (pl *PersonList) appendPerson(p Person) {
  *pl = append(*pl, p)
}
```

### Use Problem Domain Names

Problems that has specific domains should be explicitly at names. At least the programmers have a chance to ask to domain expert.

**Bad:**

```go
var t boolean
```

**Good:**

```go
var freightTracker boolean
```

### Add Meaningful Context

It's better to add context to variable names if they are from same concept.

**Bad:**

```go
var impact int
var outcome int
```

**Good:**

```go
var gameDecisionImpact int
var gameDecisionOutcome int
```

Also, we clearly can see that they can be part of a GameDecision class.

### Dont Add Gratuitous Context

Don't add context unnecessarily. With modern tools, we can easily found the classes, so it may not interesting to add some prefixes.

**Bad:**

```go
var accountAddress AccountAddress
var customerAddress CustomerAddress
var companyAddress Company Address
```

**Good:**

```go
var accountAddress Address
var customerAddress Address
var companyAddress Address
```

Obs: there is a high change that those Address classes has same attributes, so it not makes sense to create other similar classes.

## Functions

### Small

Smaller functions implies into better understand and easy to test.

**Bad:**

```go
func (cc *CompanyController) CreateCompany(c *gin.Context) {
  var companyAuth struct {
    Email    string `json:"email" binding:"required" gorm:"-"`
    Phone    string `json:"phone" binding:"required" gorm:"-"`
    Password string `json:"password" binding:"required" gorm:"-"`
    Name     string `json:"name" binding:"required" gorm:"-"`
  }
  var company model.Company

  err := c.ShouldBindBodyWith(&companyAuth, binding.JSON)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }

  err = c.ShouldBindBodyWith(&company, binding.JSON)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }

  authParams := (&auth.UserToCreate{}).
    Email(companyAuth.Email).
    PhoneNumber(companyAuth.Phone).
    Password(companyAuth.Password).
    DisplayName(companyAuth.Name)

  u, err := firebase.Conn.Auth.CreateUser(context.Background(), authParams)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }

  company.ID = u.UID

  err = DB.Conn.Create(&company).Error
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }

  c.JSON(http.StatusOK, gin.H{"status": "Success!"})
}
```

**Good:**

```go
func (cc *CompanyController) CreateCompany(c *gin.Context) {
  id, err := firebase.auth.signup(c) // here we bind the JSON as we need and do the signup
  if err != nil {
    // inside of signup, we can put the return error
    return
  }

  err = DB.Conn.Create(c, id)
  if err != nil {
    return
  }

  c.JSON(http.StatusOK, gin.H{"status": "Success!"})
}
```

### Blocks and Indenting

Probably if, else, while, etc. statements should be one line long.

**Bad:**

```go
if minimumPrice != "" {
  i, _ := strconv.Atoi(minimumPrice)
  // ...
  query = query.Where("price >= ?", i)
}
```

**Good:**

```go
if minimumPrice != "" {
  addMinimumPriceQuery(minimumPrice)
}
```

### Do one thing

Functions should do one thing. They should do it well.
> They should do it only - Jennifer M Kohnks

Using the same example of [Small](#small):

- Now the function is small
- And also only does one thing, register the user (before of that was doing some validations and bindings)

### One Level of Abstraction per Function

A function should have only one abstraction concept, it will make easier to other programmers to understand.

Using the same example of [Small](#small):

We can clearly see that the code was doing different abstract level actions, it:

- Creates the necessary structures through low level syntaxes
- Does the json binding through framework-specific function
- Access Auth Service for user signup
- Access Database Service for save user

Right at the **Good** it does as actions, but at mean functions

### Switch Statements

Avoid.

1. Each new case, it will grow the function (impacting small principle).
2. Does more than one thing.
3. Violates Single Responsability Principle.
4. Violates Open Closed Principle because it must change whenever new case is added.

**Bad:**

```go
switch orderBy {
  case "newest":
    query = query.Order("created_at DESC")

  case "highest_price":
    query = query.Order("price DESC")

  case "lowest_weight":
    query = query.Order("weight ASC")

  default:
}
```

**Good:**

```go
orderOptions := map[string]string{
  "newest":        "created_at DESC",
  "highest_price": "price DESC",
  "lowest_weight": "weight DESC",
}

query = query.Order(orderOptions[orderBy])
```

### Use Descriptive Names

You know you are working on clean code when each routine turns out to be pretty much what you expected. Half the battle is good names for small functions that do one thing.

Long descriptive name > short enigmatic name

**Bad:**

```go
// In a game context, what should be restarted? The entire game? A specific phase?
func restart() {}
```

```go
func restartGame() {}
func restartPhase() {}
func goToCheckpoint() {}
```

### Function Arguments

Ideally 0 (niladic) > 1 (monadic) > 2 (dyadic) > 3 (triadic, avoid) > polyadic (requires very special justification).
It envolves test environment too, its easy to test with 0 or 1 args, but more than that, we have a combination at tests too.

**Bad:**

```go
// To test it, will require many input permutations
func jumpStory(
  jumpStep int,
  outcome string,
  impact int,
  move int,
  jumpChapter int,
  speedBonus int,
) {}
```

```go
func jumpToPhase(jumpStep int, jumpChapter int) {}
func calculateResult(outcome int, impact int) {}
func calculateMovement(move int, speedBonus int) {}
```

### Flag Arguments

Avoid.
It immediatly proclaim that this function does more than one thing.

**Bad:**

```go
func renderGameOver(isFailure boolean) {
  if isFailure == true {
    // ...
    return render("It wasnt this time, Link")
  } else {
    // ...
    return render("Hello Zelda, goodbye Ganon")
  }
}
```

**Good:**

```go
func renderSuccessEnd() {}
func renderFailureEnd() {}
```

### Argument Objects

When a function need more than two or three args, probably could be wrapped into a class.

**Bad:**

```go
func calculateRectangleArea(height int, width, int) {}
```

**Good:**

```go
func calculateRectangleArea(rectangle Rectangle) {}

// Also, we can create a class for that:
type Rectangle struct {
  width int
  height int
}

func (r *Rectangle) area() {}
```

### Argument Lists

First, giving a function signature example:

```go
func Println(a ...interface{}) {}
```

Here, it's applied the same rules of [Function Arguments](#function-arguments) and say that its a **monadic** because we can clearly see a list of variable type as the first argument.

Other examples:

```go
func Printf(format string, a ...interface{})) {} // dyadic
func Sscanf(str string, format string, a ...interface{}) {} // triadic
```

### Verbs and Keywords

A function is an action, so it should be a verb.
Keywords at funcion name can help us to avoid cognitive process by remembering what the the function (verb) does.
Also, it mitigates the problem of having to remember the ordering of the arguments.

**Bad:**

```go
func write(name string) {}
func assertStringEquals(expect string, actual string) {}
```

**Good:**

```go
func writeField(name string) {}
func assertStringExpectedEqualsActual (expect string, actual string) {}
```

### Have no Side Effects

Side effects are lies. It goes counter of "make one thing".

**Bad:**

```go
func addStreetLine(oldLine string, lineToAdd string, fullAddress *string) string {
  newLine := oldLine + "\n" + lineToAdd
  *fullAddress = strings.Replace(*fullAddress, oldLine, newLine, 1)

  return newLine
}
```

**Good:**

```go
func addStreetLine(oldLine string, lineToAdd string) string {
  return oldLine + "\n" + lineToAdd
}
func getFullAddress() {}
```

### Output Arguments

In general output arguments should be avoided.
If your function must change the state of something, have it change the state of its owning object.

**Bad:**

```go
func addStreetLine(oldStreetLine *string, newStreetLine string) {
	*oldStreetLine += "\n" + newStreetLine
}
```

**Good:**

```go
type Address struct { /* ... */ }
func (a *Address) addStreetLine(line string) {
  a.streetLine += "\n" + line
}
```

### Command Query Separation

Functions should either do something or answer something, but not both.

**Bad:**

```go
type Anonymap map[string]string
func (a Anonymap) setAttribute(key string, value string) bool {
	if a[key] != "" {
		a[key] = value
		return true
	}
	return false
}

if a.setAttribute("name", "Yuri") == true {
  // ...
}
```

**Good:**

```go
type Anonymap map[string]string

func (a Anonymap) setAttribute(key string, value string) {
  a[key] = value
}
func (a Anonymap) attributeExists(key string) bool {
  if (a[key] != "") {
    return true
  }
  return false
}

if a.attributeExists("name") == true {
  a.setAttribute("name", "Yuri")
  // ...
}
```

### Prefer Exceptions to Returning Error Codes | Extract Try/Catch Blocks

Go doesn't have exceptions, the language's creators believe that coupling exceptions to a control structure, like `try-catch-finally` idiom, results in convoluted code.

More on [Why does Go not have exceptions?](https://golang.org/doc/faq#exceptions)

### Error Handling is One Thing

If we create another function to handle the error, we guarantee that it does one thing and not violates [Do One Thing](#do-one-thing).

**Bad:**

```go
func createUser(c *gin.Context) {
  err = Database.Create(&user).Error
	if err != nil {
    BugTracking.notify()
    authErr := firebase.DeleteUser(user.UID)

    if authErr != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": "Error on Delete User at Firebase Auth"})
      return
    }

    c.JSON(http.StatusInternalServerError, gin.H{"error": "Error on Create User at Database"})
    return
  }
}
```

We can see that the external error handling is doing three things:
- Notifying the bug tracking
- Deleting user from auth service (and also error handling it)
- Appending the error to the response

**Good:**

```go
func createUser(c *gin.Context) {
  // ...
  err = Database.Create(&user).Error
	if err != nil {
    dbCreateUserErrorHandling(user, c)
    return
	}
}

func dbCreateUserErrorHandling(user User, c *gin.Context) {
  BugTracking.notify()
  deleteUserFromAuth(user.UID)
  c.JSON(http.StatusInternalServerError, gin.H{"error": "Error on Create User at Database"})
  // Here, we can create another function for this error to not violate the [One Level of Abstraction per Function](#one-level-of-abstraction-per-function) rule.
}
```

### Dont Repeat Yourself (DRY)

This principle states that duplication in logic should be eliminated via abstraction.

It will avoid duplicate test and enforces the use of small functions.

**Bad:**

```go
func main() {
  user1 := User{
    firstName: "Yuri",
    lastName: "Kobashigawa",
  }

  user2 := User{
    firstName: "Marcelo",
    lastName: "Barreto",
  }

  user3 := User{
    firstName: "Andrew",
    lastName: "Hanasiro",
  }

  fmt.Println(user1.firstName + " " + user1.lastName)
  fmt.Println(user2.firstName + " " + user2.lastName)
  fmt.Println(user3.firstName + " " + user3.lastName)
}
```

**Good:**

```go
func main() {
	users := []User{
		User{
			firstName: "Yuri",
			lastName:  "Kobashigawa",
		},
		User{
			firstName: "Marcelo",
			lastName:  "Barreto",
		},
		User{
			firstName: "Andrew",
			lastName:  "Hanasiro",
		},
  }
  
  printUsers(users)
}

func printUsers(users []User) {
  for i := 0; i < len(users); i++ {
    fmt.Println(users[i].firstName + " " + users[i].lastName)
  }
}
```

### Structured Programming

Following Edsger Dijkstra's rules, every function and every block within, should have one entry and one exit, so only one return statement.

Avoid break, continue and never use goto statements.

5.Formatting

The code should be nicely formatted. For that we have automated tools with rules that need to be complied to the entire team.

### The Purpose of Formatting

Code formatting is about communication. It affects maintainability and extensibility.

### Vertical Formatting

Unclebob show that its possible do build significant systems with typically 200 lines long, with an upper limit of 500.

Small files are usually easier to understand than large files are.

### The Newspaper Metaphor

At the top, you should tell what the story is about
First paragraph gives you a synopsis of the whole story
At downward, the details increase

### Vertical Openness Between Concepts

All code is read left to right and top to bottom. Each line represents an expression or a clause, and each group of lines represents a complete thought. Those thoughts should be separated from each other with blank lines

### Vertical Density

If opennes separates concepts, then vertical density implies close association

### Vertical Distance

Conpects that are closely related should be kept verticall close to each other
Variable declarations should be declared as close to their usage as possible
Instance variables should be declared at the top of the class
Dependent functions should be vertically close and the caller should be above the callee, if at all possible. This gives the program a natural flow.
Conceptual affinity should be close

### Horizontal Formatting

Unclebob recommends the limit of 120 characters.
Personally I'd like to set a small font that allows more characters, but agree with him that it should not depend upon the user or the device

### Horizontal Openness and Density

White space to associate things that are strongly related and disassociate things that are more weakly related
E.g. int lineSize = line.Length()
Left side x Right side representing assignment operators
Be careful because modern formatting code tools are blind for those statements

### Horizontal Alignment

Avoid. Even putting a glance at code, the programmer get tempted to view the variable name without checking type. Also formatting tools could eliminate this kind of alignment.

So, if you use alignment because of list length, the problem is on the list that probably could be split up in separated into different classes.

### Indentation

Extremally important to visually understand hierarchy of block's information. In moderns IDEs and idiomatic languages like Go, we have tools that already help us.

### Team Rules

Every programmer has his own favorite formatting rules, but all of the team need to agree and stick with one. This will give a consistent and smooth style.

6.Objects and Data Structures

There's a reason to keep variables private, avoid dependancy of other modules and isolate the implementation

### Data Abstraction

(Create examples here)
Abstractions/interfaces hides the implementation and puts a layer of functions between the variables. Classes pushes their variables through getters and setters. Working together allows the users to manipulate the essence of the data without having to know the implementation

### Data/Object Anti-Symmetry

Object hide their data behind abstractions and expose functions that operate on that data.
Data structure expose their data and have no meaningful functions.
They are virtual opposites.

Procedural code (using data structures)) makes it easy to add new functions without changing the existing data structures. OO code, on the other hand, makes it easy to add new classes without changing existing functions.

Complement:
Procedural code makes it hard to add new data structures because all the functions must change. OO code makes it hard to add new functions because all the classes must change.

Mature programmers know that the idea that everything is an object is a myth. Sometimes you really do want simple data structures with procedures operating on them.

### The Law of Demeter

A module should not know about the innards of the objects it manipulates. Objects hide their data and expose operations. Not applied to data structures which naturally exposes their internal structure

A method f of a class C should only call the methods of these:
An object created by f
An object passed as an argument to f

An object held in an instance variable of C

In other words, talk to friends, not to strangers.

Bad example:
final String outputDir = ctxt.getOptions().getScratchDir().getAbsolutePath()

if the code had been written as data structure, its not a Demeter violations:
final String outputDir = ctxt.options.scratchDir.absolutePath

### Train Wrecks

Its called like this because it is a chain of calls that have a sloppy style and should be avoided.

Using the example above:

Options opts = ctxt.getOptions()
File scratchDir = opts.getScratchDir()
final String outputDir = scratchDir.getAbsolutePath()

### Hybrids

This confusion about what is what leads to unfortunate hybrid structure that are half object and half data structure.

### Hiding Structure

if ctxt, scratchDir and output are objects with real behavior, we should hide their access.

### Data Transfer Objects

The quintessential form of a data structure is a class with public variables and no functions. Sometimes called a data transfer object or DTO. They are very useful especially when communicating with databases or parsing messages from sockets, and so on.

### Active Record

Are special forms of DTOs. They are data structures with public variables, but typically have navigational methods like save and find. Typically are direct translations from database tables or other data sources.
Avoid to put business rules methods in them because it creates a hybrid structure.

### Conclusion

Objects expose behavior and hide data.

+ Easy to add new objects without changing existing behavior
- Harder to add new behavior to existing objects

Data structures expose data and have no significant behavior.

+ Easier to add new behavior to existing data structures
- Harder to add new data structures to existing functions

7.Error Handling

Needs to be a natural step of programing
Inputs can be unexpected abnormal and devices can fail. Programmers are responsible to create a code that do what needs to do
Error handling is important, but if it obscures logic, it's wrong

### Use Exceptions Rather Than Return Codes

In old ages, it was necessary the use of Return Codes, but it causes a chainning of statements, memory efforts to remember about some error codes and big functions
Now, we have exceptions that could error handling with smaller Catch functions, separating concerns and getting pluses on indenting/formatting and aesthetics.

### Write Your try-catch-finally Statement First

try  blocks are like transactions
catch has to leave your program in a consistent state, no matter what happens in try
Try to write tests that force exceptions, and then add behavior to your handler to satisfy your tests

### Use Unchecked Exceptions

The price of checked exceptions is an Open/Closed Principle violation. If you throw a checked exception from a method in your code and the catch is three levels above, you must declare that exception in the signature of each method between you and the catch
Checked exceptions can sometimes be useful if you are writing a critical library: you must catch them. But in general application development, the dependency costs outweigh the benefits

### Provide Context with Exceptions

Each exception should provie enough context to determine the source and location of an error. Create informative error messages and pass them along with your exceptions. Mention the operation that failed and the type of failure

### Define Exception Classes in Terms of a Caller's Needs

There are many ways to classify errors. The most important concern should be how they are caught
Wrapping third-party API minimize your dependencies upon it, if you choose to move to a different library in the future, it will be without much penalty. Also makes it easier to mock out when testing

### Define the Normal Flow

Following the preceding sections, you'll end up with a good amount of separation between your business logic and your error handling
But there are cases that we can use mistakely a error handling coupled with our business logic without the need.

### Don't Return Null

If you are tempted to return null from a method, consider throwing and exception or returning a SPECIAL CASE object instead

### Don't Pass Null

Avoid passing null in your code whenever possible

### Conclusion

Clean code is readable, but it must also be robust. These are not conflicting goals. We can write robust clean code if we see error handling as a separate concern, something that is viewable independently of our main logic. To the degree that we are able to do that, we can reason about it independently, and we can make great strides in the maintainability of our code.

8.Boundaries
Sometimes we buy third-party packages, use open source or depend on teams in own company to produce components or subsystems for us. Somehow we must cleanly integrate this foreign code with our own. This chapter tells good practices and techniques to keep the boundaries of our software clean.

### Using Third-Party Code

Unclebob uses Map (Java) as example
If the application needs only some behavior from Map without the need of clear() method, an incorrect use could bring this around the entire application
Also, at Java, programmers could need to cast their type or use generics around the code
The best way to use a Third Party Code with specific behaviors and type is by wrapping the code in Classes, keeping the TPC inside the classe or close family of classes

### Exploring and Learning Boundaries

Third-party code helps us get more functionality delivered in less time. But learning it is hard, integrating it is hard too. Doing both at the same time is doubly hard
Instead of experimenting and trying out the new stuff in our production code, we could write some tests to explore our understanding of the third-party code. In those learning tests we call the third-party API as we expect to use it in our application

### Learning Tests Are Better Than Free

It costs nothing. We had to learn the API anyway, and wrigint those tests was an easy and isolated way to get that knowledge. Also, when new releases occurs, we can run our tests to see if there are behavioral differences

### Using Code That Does Not Yet Exist

If we encounter a project phase that we need to use a code that haven't been released yet and we don't want to being blocked, we can create a fake interface describing some behaviors that we expect. After the release, we can integrate it using a design pattern called THE ADAPTER to encapsulate the interaction with the API and provide a single place to change it when the API evolves
This design provides a convenient stub to test the behaviors.

### Clean Boundaries

When we use code that is out of our control, special care must be taken to protect our investment and make sure future change is not too costly
Code at the boundaries needs clear separation and tests that define expectations. We should avoid letting toomuch of our code know about the third-party particulars. It's better to depend on something you control than on something you don't control, lest it end up controlling you

9.Unit Tests

Unclebob quotes one experience of him. As summary, nowadays he would write tests that make sure any piece of code works as expected. Including stubs and mocks. After that, he would improve the suit of tests to run on every environment to make sure that his co-workers could run.
Agile and TDD movements have encouraged programmers to write automated unit tests.

### The Three Laws of TDD

Write unit tests first, before write production code is just the tip of the iceberg, here are the three laws

First Law: You may not write production code until you have written a failing unit test
Second Law: You may not write more of a unit test than is sufficient to fail, and not compiling is failing
Third Law: You may not write more production code than is sufficient to pass the currently failing test

If we work this way, we will write dozens of tests every day, hundreds of tests every month and thousands of tests every year, covering virtually all of our production code

### Keeping Tests Clean

Tests must change as the production code evolves. The dirtier the tests, the harder they are 

to change. If the variables aren't well named, functions aren't short and descriptive and test code aren't well designed and thoughtfully partitioned, its a dirty test

If it is accumulative, can turn into a snowball that elevates the estimative of functionalities because of maintaing tests. Raising and raising, test suite will be dropped. losting the ability to make sure that changes to their code base worked as expected

Test code is just as important as production code, must be kept as clean as production code

### Tests Enable the -ilities

It is unit tests that keep our code flexibe, maintainable and reusable. The reason is simple. If you have tests, you don't fear making changes to the code. Without tests every change is a possible bug

### Clean Tests

Readability: clarity, simplity and density of expression
Unclebob uses an example of FitNesse by comparing a v1 with lots of duplication of code and functions with abstraction level that are irevelant for the reader. At v2, he uses a BUILD-OPERATE-CHECK pattern and also going right to the point, using only data types and functions that truly need, easing to readers that could quickly understand what it does without being misled or overwhelmed by details

### Domain-Specific Testing Language

Rather than using the APIs that programmers use to manipulate the systems, its better build a set of functions and utilities that make use of those APIs and that make the tests more convenient to write and easier to read. These functions and utilities become a specialized API used by the tests. They are a testing language

### A Dual Standard

There are things that you might never do in a production environment that are perfectly fine in a test environment. Usually they involve issues of memory or CPU efficiency. But they never involve issues of cleanliness

### One Assert per Test

There are sometimes that we can avoid more than one assert even being part of the same context. It will result in duplication of code that, as consequence of cleanliness, it will have some minor refactors and some patterns that could be done, but it shows a high price to separate things that are at same concept.
As summary, best thing we can say is that the number of asserts in a test ought to be minimized

### Single Conpect per Test

Perphaps a better rule is that we sant to test a single concept in each test function, as 

explained above

### F.I.R.S.T.

Clean tests follow five other rules:
F.ast: Tests should be fast. When tests run slow, you won't want to run them frequently
I.ndependent: Tests should not depend on each other. When tests depend on each other, then the first one to fail causes a cascade of downstream failures, making diagnosis difficult and hiding downstream defects
R.epeatable: Tests should be repeatable in any environment, Production, Staging, QA, your laptop while riding home on the train without a network
S.elf-validating: Tests should have a boolean output. Either they pass or fail. You should not have to read a log file to tell the situation. If the tests aren't self-validating, then failure can become subjective and running the tests can require a long manual evaluation
T.imely: Tests need to be written in a timely fashion. Unit tests should be written just before the production code that makes them pass. If you write tests after the production code, then you may find the production code to be hard to test

### Conclusion

Tests are as important to the health of a project as the production code is. Perhaps they are even more important, because tests preserve and enhance the flexibility, maintainability and reusability of the production code
If you let the tests rot, then your code will rot too. Keep your tests clean

10.Classes

### Class Organization

Following Java convention:
Class should begin with a list of variables
Public static constants, if have, should come first
Private static variables
Private instance variables
Public variable (if have a very good reason)
Public functions
Private utilities called by public function right after the public function itself
This follows the stepdown rule and helps the program read like a newspaper article

### Encapsulation

Should keep variables and utility functions private, but sometimes could make it protected so that can be accessed by a test. Remember that loosening encapsulation is always a last resort

### Classes Should be Small!

Functions are measured size by counting physical lines
Classes are measured by responsibilities

The name of a class should describe what responsibilities it fulfills, probably the first way of helping determine class size. The more ambiguous the class name, the more likely it  has too many responsibilities

We should also be able to write a brief description of the class in about 25 words without using words like "if", "and", "or", or "but"

### The Single Responsibility Principle (SRP)

A class or module should have one, and only one, reason to change.
Trying to identify responsibilities (reasons to change) often helps us recognize and create better abstractions in our code

SRP is one of the more important concept in OO design. Also one of the simpler concepts to understand and adhere to. But regularly we find classes that do many things. Why?

Part of developers thought that when the code works, the job is done. They don't think that organization and cleanliness is part of the job
Part of developers thought that having lots of small classes could be hard to navigate in order to figure how it works. But this navigation brings a learning flow into code

Every sizable system will contain a large amount of logic and complexity. The primary goal in managing such complexity is to organize it so that a developer knows where to look to find things and need only understand the directly affected complexity at any given time.
We want our systems to be composed of many small classes, not a few large ones. Each small class encapsulates a single responsibility, has a single reason to change, and collaborates with a few others to achieve the desired system behaviors

### Cohesion

The more variables a method manipulates the more cohesive that method is to its class. A class in which each variable is used by each method is maximally cohesive. When cohesion is high, it means that the methods and variables of the class are co-dependent and hang together as a logical whole

The strategy of keeping functions small and keeping parameter lists short can sometimes lead to a proliferation of instance variables that are used by a subset of methods. When this happens, it almost always means that there is at least one other class trying to get out of the larger class. You should try to separate the variables and methods into two or more classes such that the new classes are more cohesive.

### Maintaining Cohesion Results in Many Small Classes

Small classes is a natural flow when you maintain cohesion because you break large function into many smaller functions and see the opportunity to recognize and split the big class into smaller ones

### Organizing for Change

For most systems, change is continual. Every change subjects us to the risk that the remainder of the system no longer works as intended. In a clean system we organize our classes so as to reduce the risk of change

Private method behavior that applies only to a small subset of a class can be a useful heuristic for spotting potential areas for improvement

### Isolating from change

Needs will change, therefore code will change
A client class depending upon concrete details is at risk when those details change. We can introduce interfaces and abstract classes to help isolate the impact of those details

The lack of coupling means that the elements of our system are better isolated from each other and from change. This isolation makes it easier to understand each element of the system

Unclebob uses Portfolio example that depends upon concrete class TokyoStockExchange. He creates the StockExchange abstract class and now Tokyo... implements this abstract class. Also Portfolio pass to depend upon an abstract class, decoupling Portfolio from all specific details of a specific X Stock such as price or format
By minimizing coupling in this way, our classes adhere to another class design principle known as the Dependency Inversion Principle (DIP) that says that our classes should depend upon abstractions, not on concrete details

11.Systems
"Complexity kills. It sucks the life out of developers, it makes products difficult to pan, build, and test." - Ray Ozzie, CTO, Microsoft Corporation

### How Would you Build a City?

Can't manage all the details alone. Cities works because there are teams of people who manage particular parts. Also work because they have evolved appropriate levels of abstraction and modularity that make it possible for individuals and the "components" they manage to work effectively, even without understanding the big picture

Here we will consider how to stay clean at higher levels of abstraction, the system level

### Separate Constructing a System from Using It

Consider that construction is a very different process from use

Software systems should separate the startup process, when the application objects are constructed and the dependencies are "wired" together, from the runtime logic that takes over after startup

We should modularize the process of object construction separately from the normal runtime logic and we should make sure that we have a global, consistent strategy for resolving our major dependencies

### Separation of Main

One way to separate construction from use is simply to move all aspects of construction to main, or modules called by main, and to design the rest of the system assuming that all objects have been constructed and wired up appropriately

The flow of control is easy to follow. main function builds the objects necessary for the system, then passes them to the application, which simply uses them

### Factories

Sometimes, we need to make the application responsible for when an object gets created. In this case we can use the ABSTRACT FACTORY pattern to give the application control of when to build, but keep the details of that construction separate from the application code

### Dependency Injection

A powerful mechanism for separating construction from use is Dependency Injection (DI), the application of Inversion of Control(IoC) to dependency management. Inversion of Control moves secondary responsibilities from an object to other objects that are dedicated to the purpose, thereby supporting the Single Responsibility Principle. In the context of dependency management, an object should not take responsibility for instantiating dependencies itself. Instead, it should pass this responsibility to another "authoritative" mechanism, thereby inverting the control

The invoking object doesn't control what kind of object is actually returned (as long it implements the appropriate interface), but the invoking object still actively resolves the dependency

True Dependency Injection goes one step further. The class takes no direct steps to resolve its dependencies; completely passive. Instead, it provides setter methods or constructor arguments (or both) that are used to inject the dependencies

Most DI containers won't construct an object until needed (LAZY-INITIALIZATION). Also, many of these containers provide mechanisms for invoking factories or fo constructing proxiess, which could be used for LAZY-EVALUATION and similar optimizations

### Scaling Up

Its a myth that we can get systems "right the first time". Instead, we should implement only today's stories, then refactor and expand the system to implement new stories tomorrow. This is the essence of iterative and incremental agility. TDD, refactoring, and the clean code they produce make this work at the code level

And System Level?

Software systems are unique compared to physical systems. Their architectures can grow incrementally, if we maintain the proper separation of concerns

### Cross-Cutting Concerns

When you have to spread essentially the same code that implements the persistence strategy across many objects
E.g. persistence, transactions, security, caching, failover

### Test Drive the System Architecture

An optimal system architecture consists of modularized domains of concern, each of which is implemented with Plain Old Java (or other) Objects. The different domains are integrated together with minimally invasive Aspects or Aspect-like tools. This architecture can be test-driven, just like the code

### Optimize Decision Making

Modularity and separation of concerns make decentralized management and decision making possible. In a sufficiently large system, whether it is a city or a softwware project, no one person can make all the decisions

Its best to give responsibilities to the most qualified person. Also best to postone decisions until the last possible moment, it lets us make informed choices with the best possible information. A premature decision is a decision made with suboptimal knowledge

The agility provided by a POJO system with modularized concerns allows us to make optimal, just-in-time decisions, based on the most recent knowledge. The complexity of these decisions is also reduced

### Use Standards Wisely, When They Add Demonstrable Value

Standards make it easier to reuse ideas and components, recruit people with relevant experience, encapsulate good ideas, and wire components together. However, the process.

of creating standards can sometimes take too long for industry to wait, and some standards lose touch with the real needs of the adapters they are intended to serve

### Systems Need Domain-Specific Languages

A good Domain-Specific Language (DSL) minimizes the "communication gap" between a domain concept and the code that implements it. If you are implementing domain logic in the same language that a domain expert uses, there is less risk that you will incorrectly translate the domain into the implementation

DSLs, when used effectively, raise the abstraction level above code idioms and design patterns. They allow the developer to reveal the intent of the code at appropriate level of abstraction

Domain-Specific Languages allow all levels of abstraction and all domains in the application to be expressed as POJOs, from high-level policy to low-level details

### Conclusion

System must be clean too. An invasive architecture overwhelms the domain logic and impacts agility. When the domain logic is obscured, quality suffers because bugs find it easier to hide and stories become harder to implement. If agility is compromised, productivity suffers and the benefits of TDD are lost

12.Emergence

### Getting Clean via Emergent Design

These rules making easier to apply principles such as SRP and DIP, facilitating the emergence of good designs

According to Kent, a design is "simple" if it follows these rules:
Runs all the tests
Contains no duplication
Expresses the intent of the programmer
Minimizes the number of classes and methods
The rules are given in order of importance

Simple Design Rule 1: Runs All the Tests
System that aren't testable aren't verifiable. Arguably, a system that cannot be verified should never be deployed

The more test we write, the more we'll bring things that are simpler to test, as consequence, SRP comes.

Tight couplking makes it difficult to write tests. Similarly, the more test we write, more we use principles like DIP and tools like dependency injection, interfaces and abstraction to minimize coupling. Our designs improve even more

We also achieve low coupling and high cohesion. So, Writing tests leads to better designs

### Simple Design Rules 2-4: Refactoring

After run all of tests and they pass, we eliminates the fear that cleaning up the code will break it. So we can go through a refactoring steps aplying all of our knowledge about good software design, increasing cohesion, decreasing coupling, separating concerns, modularize system concerns, choose better names, and so on. Consequently aplying the final three rules: eliminate duplication, ensure expressiveness, and minimize the number of classes and methods

### No Duplication

Duplication is the primary enemy of a well-designed system. Represents additional work, risk, unnecessary complexity
As we eliminate duplication, also its better to see possible violations of SRP

### Expressive

The clearer the author can make the code, the less time others will have to spend understanding it. This will reduce defects and shrink the cost of maintenance
You can express yourself by:
Choosing good names. A bad name could bring surprises
Keeping functions and classes small. Its easy to name, to write and to understand
Using standard nomenclature
Well-written unit tests. Once he can act as a documentation, quickly describing what a code does and what to expect

### Minimal Classes and Methods

High class and method counts are sometimes the result of pointless dogmatism. The goal is keep overall system small while also keep the functions and classes small

### Conclusion

Following the practice of simple design can and does encourage and enable developers to adhere to good principles and patterns that otherwise take years to learn

13.Concurrency

"Objects are abstractions of processing. Threads are abstractions of schedule." - James O.Coplien

Writing clean concurrent programs is very hard. Its much easier to write code that executes in single thread. Also easy to write multithreaded code that looks fine on the surface, but broken in a depper level. Such code works fine until the system is placed under stress

### Why Concurrency?

Concurrency is a decoupling strategy. Help us decouple what gets done from when it gets done. This can dramatically improve both the throughput and structures of an application. From a structural point of view the application looks like many little collaborating computers rather than one big main loop. This can make the system easier to understand and offers some powerful ways to separate concerts

E.g. Consider a system that interprets large data sets but can only give a complete solution after processing all of them. Perhaps each data set could be processed on a different computer, so that many data sets are being processed in parallel

### Myths and Misconceptions

Concurrency always improves performance - sometimes improve performance, but only when there is a lot of wait time that can be shared between multiple threads or multiple processor. Neither situation is trivial
Design does not change when writing concurrent programs - they can be remarkably different from the design of a single-threaded system. The decoupling of what from when usually has a huge effect on the structure of the system
Undestanding concurrency issues is not important when working with a container such as a Web or EJB container - better know just what your container is doing and how to guard against the issues of concurrent update and deadlock 

Here are more balanced sound:
Concurrency incurs some overhead - both in performance as well as writing additional code
Correct concurrency is complex - even for simple problems
Concurrency bugs aren't usually repeatable - so they are often ignored as one-offs instead of the true defects they are
Concurrency often requires a fundamental change in design strategy

### Concurrency Defence Principles

SRP - keep your concurrency-related code separate from other code
Corollary: Limit the Scope of Data - take data encapsulation to heart; severely limit the access of any data that may be shared
Corollary: Use Copies of Data - avoid sharing objects
Corollary:Threads Should Be as Independent as Possible - Attempt to partition data into independent subsets than can be operated on by independent threads, possibly 

in different processors
Know Your Library - Review the classes available to you. In the case of Java, become familiar with java.util.concurrent, java.util.concurrent.atomic, java.util.concurrent.locks
Know Your Execution Models
First, some important definitions:
Bound Resources - Resources of a fixed size or number used in a concurrent environment. E.g. database connections and fixed-size read/write buffers
Mutual Exclusion - Only one thread can access shared data or resource at a time
Starvation - One thread or a group is prohibited from proceeding for an excessively long time or forever. E.g. always letting fast-running threads through first could starve out longer running threads if there is no end to the fast-running threads
Deadlock - Two or more threads waiting for each other to finish. Each thread has a resource that the other thread requires and neither can finish until it gets the other resource
Livelock - Threads in lockstep, each trying to do work but finding another "in the way". Due to resonance, threads continue trying to make progress but are unable to for an excessively long time or forever

Now, the various execution models used in concurrent programming:
Producer-Consumer: One or more producer threads create some work and plate it in a buffer or queue. One or more consumer threads acquire that work from the queue and complete it. The queue between the producers and consumers is a bound resource. This means producers must wait for free space in the queue before writing and consumers must wait until there is something in the queue to consume. Coordination between the producers and consumers via the queue involves producers and consumers signaling each other. The producers write to the queue and signal that the queue is no longer empty. Consumers read from the queue and signal that the queue is no longer full. Both potentially wait to be notified when they can continue
Readers-Writers: When you have a shared resource that primarily serves as a source of information for readers, but which is occasionally updated by writers, throughput is an issue. Emphasizing thorughput can cause starvation and the accumulation of stale information. Allowing updates can impact thorughput. Coordinating readers so they do not read something a writer is updating and vice versa is a tough balancing act. Writers tend to block many readers for a long period of time, thus causing throughput issues. The challenge is to balance the needs of both readers and writers to satisfy correct operation, provide reasonable throughput and avoiding starvation. A simple strategy makes writers wait until there are no readers before allowing the writer to perform an update. If there are continuous 

readers, however, the writers will be starved. On the other hand, if there are frequent writers and they are given priority, throughput will suffer. Finding that balance and avoiding concurrent update issues is what the problem addresses
Dining Philosophers: Imagine a number of philosophers sitting around a circular table. A fork is placed to the left of each philosopher. There is a big bowl of spaghetti in the center of the table. The philosophers spend their time thinking unless they get hungry. Once hungry, they pick up the forks on either side of them and eat. A philosopher cannot eat unless he is holding two forks. If the philosopher to his right or left is already using one of the forks he needs, he must wait until that philosopher finishes eating and puts the forks back down. Once a philosopher eats, he puts bot his forks back down on the table and waits until he is hungry again. Replacing philosophers with threads and forks with resources and this problem is similar to many enterprise applications in which processes compete for resources. Unless carefully designed, systems that compete in this way can experience deadlock, livelock, throughput and efficiency degradation. Most of concurrent problems you will likely encounter will be some variation fo these three problems. Study these algorithms and write solutions using them on your own so that when you come across concurrent problems, you'll be more prepared to solve the problem
Beware Dependencies Between Synchronized Methods - Avoid using more than one method on a shared object. There will be times when you must use more than one method on a shared object, so there are three ways to make the code correct:
Client-Based Locking - Have the client lock the server before calling the first method and make sure the lock's extend includes code calling the last method
Server-Based Locking - within the server create a method that locks the server, calls all the methods, and then unlocks. Have the client call the new method
Adapted Server - create an intermediary that performs the locking. This is an example of server-based locking, where the original server cannot be changed
Keep Synchronized Sections Small
Writing Correct Shut-Down Code Is Hard - Think about shut-down early and get it working early. It's going to take longer than you expect. Review existing algorithms because this is probably harder than you think
Testing Threaded Code - does not guarantee correctness, but good testing can minimize risk. Write test that have the potential to expose problems and then run them frequently, with different programatic configurations and system configurations and load. If tests ever fail, track down the faiure. Don't ignore a failure just because the tests pass on a subsequent run

Threat Spurious Failures as Candidate Threading Issues - don't ignore them
Get Your Nonthreaded Code Working First - don't try to chase down nonthreading bugs and threading bugs at the same time. Make sure your code works outside of threads
Make Your Threaded Code Pluggable - so that you can run it in various configurations
Make Your Threaded Code Tunable - getting the right balance of threads typically requires trial and error, so allow the number of threads to be easily tuned
Run with More Threads Than Processors - Things happen when the system switches between tasks, the more frequently your tasks swap, the more likely you'll encounter code that is missing a critical section or causes deadlock
Run on Different Platforms - all target platforms early and often
Instrument Your Code to Try and Force Failures - its normal for flaws in concurrent code to hide, so use jiggling strategies to ferret out errors

### Conclusion
Concurrent code is difficult to get right. First and foremost, follow the Single Responsibility Principle.
