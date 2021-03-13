class Man(object):

    def __init__(self,name,age):
        self.name = name
        self.age = age

    def walk(self):
        print("%s is walking"%self.name)
    def sing(self):
        print("%s is singing"%self.name)
    def info(self):
        print("%s is %d years old"%(self.name,self.age))

man = Man("张三",16)
man.walk()
man.sing()
man.info()