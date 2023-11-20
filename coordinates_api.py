from flask import Flask, request, jsonify
from geopy.geocoders import Nominatim

app = Flask(__name__)

@app.route('/get_coordinates', methods=['POST'])
def get_coordinates():
    data = request.get_json()

    cep = data.get('cep')
    street = data.get('street')
    neighborhood = data.get('neighborhood')
    city = data.get('city')
    state = data.get('state')

    address = f'{street} {neighborhood} {city} {state}'

    geolocator = Nominatim(user_agent='geoapi')
    location = geolocator.geocode(cep) or geolocator.geocode(address)

    try:
        coordinates = {
            'latitude': location.latitude,
            'longitude': location.longitude
        }

        return jsonify(coordinates)

    except AttributeError as e:
        return jsonify({'error': f'Error accessing latitude and longitude: {e}'})

    except Exception as e:
        return jsonify({'error': f'An unexpected error occurred: {e}'})

if __name__ == '__main__':
    app.run(debug=True)
