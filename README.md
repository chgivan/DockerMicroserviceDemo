# DockerMicroserviceExample

### Περιγραφή
Οι χρήστες  της εφαρμόγης μπόρουν να αφήσουν μηνύματα στην ιστοσελίδα και να κάνουν "Like" σε άλλα μηνύματα.

### Σκοπός
Σκοπός της εφαρμόγης είναι να δείξει την υλοποιήση μιάς Microservice Αρχιτεκτονικής, στην πλατφόρμα Docker.

### Οδηγίες εγκατάστασης
Έναρξη της εφαρμογής.
```sh
docker-compose build
docker-compose up
```

Προσωρινή διακοπή της εφαρμογής.
```sh
docker-compose kill
```

Διαγραφή των εκτελέσιμων Containers της εφαρμογής.
```sh
docker-compose down --volumes
``` 