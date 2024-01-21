from flask import Flask, request, jsonify
from geopy.geocoders import Nominatim

app = Flask(__name__)

@app.route('/get_coordinates', methods=['GET'])
def get_coordinates():
    geolocator = Nominatim(user_agent='geoapi')
    
    data = request.get_json()
    zip_code = data.get('zip_code')

    if zip_code:
        location = geolocator.geocode(zip_code)

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
