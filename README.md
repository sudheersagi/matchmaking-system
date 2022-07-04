# matchmaking-system

This project is build on open-match framework.This is the customized version of matchfunction and its dependent modules.



# Run the below steps in each module folder to set up the Match Function.

Step1: Specify your DockerHub URL.
```
REGISTRY=[YOUR_DOCKERHUB_URL]
```

Step2: Build the Match Function image.
```
docker build -t $REGISTRY/matchmaking-system-matchfunction .
```

Step3: Push the Match Function image to the configured Registry.
```
docker push $REGISTRY/matchmaking-system-matchfunction
```

Follow similarly for the other modules.
