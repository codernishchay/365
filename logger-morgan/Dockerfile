From node:alpine 

WORKDIR =  /app         
COPY package.json ./               
RUN npm i                           
COPY ./ ./     

CMD ["npm" , "start"]

# // create app in container
#   // copy package.json to container
# //install in dockercontainer
# // copy all the files into the container 