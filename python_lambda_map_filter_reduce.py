
# coding: utf-8

# # Lambda , Map , Filter and Reduce

# ## Lambda

# In[ ]:


sum = lambda x, y : x + y


# In[ ]:


print(sum(5,5))


# ## Map

# In[ ]:


#####In python2
#returns a list object.
#So it is slow and uses more memory

####In python3
#returns a generator/lazy function and they dont produce results until the values are consumed.
#So it is faster and uses less memory


# In[ ]:


mylist = [1,2,3,4,5]


# In[ ]:


result = map(lambda x : x * 2, mylist)


# In[ ]:


print(result)


# In[ ]:


print(type(result))


# In[ ]:


print(list(result))


# ###### OR

# In[ ]:


mylist = [1,2,3,4,5]


# In[ ]:


myfunc = lambda x : x * 2


# In[ ]:


result = map(myfunc, mylist)


# In[ ]:


print(result)


# In[ ]:


print(type(result))


# In[ ]:


print(list(result))


# ###### OR

# In[ ]:


mylist = [1,2,3,4,5]


# In[ ]:


def myfunc(x):
    return x * 2


# In[ ]:


result = map(myfunc, mylist)


# In[ ]:


print(result)


# In[ ]:


print(type(result))


# In[ ]:


print(list(result))


# ## Filter

# In[ ]:


#####In python2
#returns a list object.
#So it is slow and uses more memory

####In python3
#returns a generator/lazy function and they dont produce results until the values are consumed.
#So it is faster and uses less memory


# In[ ]:


mylist = [1,2,3,4,5]


# In[ ]:


result = filter(lambda x : x % 2 == 0, mylist)


# In[ ]:


print(result)


# In[ ]:


print(type(result))


# In[ ]:


print(list(result))


# ###### OR

# In[ ]:


mylist = [1,2,3,4,5]


# In[ ]:


def myfunc(x):
    return x * 2


# In[ ]:


result = filter(myfunc, mylist)


# In[ ]:


print(result)


# In[ ]:


print(type(result))


# In[ ]:


print(list(result))


# ## Reduce

# In[ ]:


#####In python2
#Built in fuction

####In python3
#moved to functools modules


# In[ ]:


import functools


# In[ ]:


mylist = [1,2,3,4,5]


# In[ ]:


result = functools.reduce(lambda x, y: x + y, mylist)


# In[ ]:


print(result)


# In[ ]:


print(type(result))


# In[ ]:




