"""create digital contents

Revision ID: ed0bc7456050
Revises:
Create Date: 2019-05-16 16:18:42.969682

"""
import sqlalchemy as sa
from alembic import op

# revision identifiers, used by Alembic.
revision = 'ed0bc7456050'
down_revision = None
branch_labels = None
depends_on = None


def upgrade():
    """
    media_id is uuid generate by time
    category_id is category table
    api_key_id is from api_gate_way generate and it's in companies table also
    publisher_id is who is sending data to our server, basically it's in publisher table
    """
    op.create_table(
        "digital_contents", sa.Column('id', sa.Integer, primary_key=True),
        sa.Column('api_key_id', sa.String(255), nullable=False),
        sa.Column('media_id', sa.String(255), nullable=False),
        sa.Column('category_id', sa.INTEGER, nullable=False),
        sa.Column('publisher_id', sa.INTEGER, nullable=False),
        sa.Column('digest', sa.Text(), nullable=False),
        sa.Column('sha256', sa.String(255), nullable=False),
        sa.Column('size_file', sa.String(255), nullable=False),
        sa.Column('title', sa.String(255), nullable=False),
        sa.Column('author', sa.String(255), nullable=False),
        sa.Column('country_of_origin', sa.String(255), nullable=False),
        sa.Column('language', sa.String(255), nullable=False),
        sa.Column('publish_date', sa.Date, nullable=False),
        sa.Column('paperback', sa.INTEGER, nullable=False),
        sa.Index("title_index", "book_id", unique=True),
        sa.Index("media_id_index", "book_id", unique=True))


def downgrade():
    op.drop_table('digital_contents')
