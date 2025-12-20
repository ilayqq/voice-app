import React from 'react';
import { StatusBar } from 'expo-status-bar';
// import AppNavigator from './src/navigation/AppNavigator';
import {SafeAreaView} from "react-native-safe-area-context";
import {StyleSheet} from "react-native";

export default function App() {
    return (
        <SafeAreaView style={styles.container}>
            {/*<AppNavigator />*/}
            <div>Hello</div>
            <StatusBar style="auto" />
        </SafeAreaView>
    );
}

const styles = StyleSheet.create({
    container: { flex: 1 }
});