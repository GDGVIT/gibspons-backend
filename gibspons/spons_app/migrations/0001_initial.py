# Generated by Django 4.2.6 on 2024-01-10 17:28

from django.db import migrations, models
import django.db.models.deletion


class Migration(migrations.Migration):

    initial = True

    dependencies = [
    ]

    operations = [
        migrations.CreateModel(
            name='Company',
            fields=[
                ('id', models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('name', models.CharField(max_length=254)),
                ('website', models.URLField()),
                ('domain', models.CharField(max_length=50)),
                ('status', models.CharField(max_length=20)),
            ],
        ),
        migrations.CreateModel(
            name='POC',
            fields=[
                ('id', models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('name', models.CharField(max_length=254)),
                ('email', models.EmailField(max_length=254)),
                ('linkedin', models.URLField()),
                ('phone', models.CharField(blank=True, max_length=15, null=True)),
                ('company', models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, to='spons_app.company')),
            ],
        ),
        migrations.CreateModel(
            name='Event',
            fields=[
                ('id', models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('name', models.CharField(max_length=255)),
                ('date_of_event', models.DateField()),
                ('expected_reg', models.IntegerField()),
                ('brochure', models.URLField()),
                ('money_raised', models.IntegerField()),
                ('created_at', models.DateTimeField(auto_now_add=True)),
                ('sponsors', models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, to='spons_app.company')),
            ],
        ),
    ]
