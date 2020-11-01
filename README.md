# goPrototypeDesignPattern

The prototype design pattern in go - The prototype pattern reflects a certain real worldliness about object and `thing` creation. In the real world it is seldom that complex objects are made completely from scratch. It is far more typical for objects to be created from reiterating existing designs, or by composing from existing and extended objects.

In this pattern we take a (partially or fully constructed) design as the prototype. We make a copy and extend it. To do this we require

- `Deep copy` support
- This cloning can be made to a factory.

Here we can see an example of the shallow copying issue and how a deep copy prevents it, bt while this code examples how we can solve the top level issue it is not a scalable solution as in a many person application it would lead to a lot of code.

![](/assets/prototype1.png)

We can add deep copy methods to do our handling for us, this is cleaner than the first solution above.

![](/assets/prototype2.png)

We can further expand on this with serlialization

![](/assets/prototype3.png)

We can then use the factory methods to generate convenience techniques for object creation.

![](/assets/prototype4.png)
